package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/domain/user"
	"github.com/booyakaasha/reddneck/internal/dto"
)

func (db *DB) CreatePost(ctx context.Context, post post.Post) error {
	query := builder().Insert("post").Columns(
		"id",
		"user_id",
		"title",
		"content",
		"created_at",
		"updated_at",
	).Values(
		post.ID,
		post.UserID,
		post.Title,
		post.Content,
		post.CreatedAt,
		post.UpdatedAt,
	).Suffix("on conflict do nothing")

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("query.ToSql: %w", err)
	}

	_, err = db.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}

func (db *DB) GetPostByID(ctx context.Context, id post.ID) (post.Post, error) {
	query := builder().Select(
		"id",
		"user_id",
		"title",
		"content",
		"created_at",
		"updated_at",
	).From("post").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return post.Post{}, fmt.Errorf("query.ToSql: %w", err)
	}

	p, err := scanPost(db.db.QueryRowContext(ctx, sql, args...))
	if err != nil {
		return post.Post{}, fmt.Errorf("scanPost: %w", err)
	}

	return p, nil
}

func (db *DB) GetPosts(ctx context.Context, crsr cursor.Cursor[*dto.GetPostsCursor]) (dto.GetPostsResult, error) {
	query := builder().Select(
		"id",
		"user_id",
		"title",
		"content",
		"created_at",
		"updated_at",
	).From("post")

	if crsr.Value != nil {
		switch crsr.Direction {
		case cursor.DirectionNext:
			query = query.Where("id <= ?", crsr.Value.ID)
		case cursor.DirectionPrev:
			query = query.Where("id > ?", crsr.Value.ID)
		}
	}

	switch crsr.Direction {
	case cursor.DirectionNext:
		query = query.OrderBy("id")
	case cursor.DirectionPrev:
		query = query.OrderBy("id")
	}
	query = query.Limit(uint64(crsr.Limit) + 1)

	sql, args, err := query.ToSql()
	if err != nil {
		return dto.GetPostsResult{}, fmt.Errorf("query.ToSql: %w", err)
	}

	rows, err := db.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return dto.GetPostsResult{}, fmt.Errorf("db.QueryContext: %w", err)
	}
	defer rows.Close()

	posts := make([]post.Post, 0, crsr.Limit+1)
	for rows.Next() {
		var post post.Post

		post, err = scanPost(rows)
		if err != nil {
			return dto.GetPostsResult{}, fmt.Errorf("scanPost: %w", err)
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return dto.GetPostsResult{}, fmt.Errorf("rows.Err: %w", err)
	}

	if len(posts) == 0 {
		return dto.GetPostsResult{}, nil
	}

	var cursorGroup cursor.Group

	if crsr.Value != nil {
		cursorGroup.Prev = dto.GetPostsCursorFromPost(posts[0]).Marshal()
	}

	if len(posts) == crsr.Limit+1 {
		cursorGroup.Next = dto.GetPostsCursorFromPost(posts[len(posts)-1]).Marshal()
		posts = posts[:len(posts)-1]
	}

	return dto.GetPostsResult{
		Posts:       posts,
		CursorGroup: cursorGroup,
	}, nil
}

func scanPost(row Row) (post.Post, error) {
	var (
		id        string
		userID    string
		title     string
		content   string
		createdAt time.Time
		updatedAt time.Time
	)

	if err := row.Scan(
		&id,
		&userID,
		&title,
		&content,
		&createdAt,
		&updatedAt,
	); err != nil {
		return post.Post{}, fmt.Errorf("rows.Scan: %w", err)
	}

	return post.Post{
		ID:        post.MustNewIDFromString(id),
		UserID:    user.MustNewIDFromString(userID),
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
