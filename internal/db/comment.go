package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/comment"
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/domain/user"
	"github.com/booyakaasha/reddneck/internal/dto"
)

func (db *DB) CreateComment(ctx context.Context, comment comment.Comment) error {
	query := builder().Insert("post").Columns(
		"id",
		"parent_id",
		"post_id",
		"user_id",
		"content",
		"created_at",
		"updated_at",
	).Values(
		comment.ID,
		comment.ParentID,
		comment.PostID,
		comment.UserID,
		comment.Content,
		comment.CreatedAt,
		comment.UpdatedAt,
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
func (db *DB) GetCommentByID(ctx context.Context, id comment.ID) (comment.Comment, error) {
	query := builder().Select(
		"id",
		"parent_id",
		"post_id",
		"user_id",
		"content",
		"created_at",
		"updated_at",
	).From("comment").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return comment.Comment{}, fmt.Errorf("query.ToSql: %w", err)
	}

	c, err := scanComment(db.db.QueryRowContext(ctx, sql, args...))
	if err != nil {
		return comment.Comment{}, fmt.Errorf("scanPost: %w", err)
	}

	return c, nil
}

func (db *DB) GetPostComments(
	ctx context.Context,
	postID post.ID,
	crsr cursor.Cursor[*dto.GetPostCommentsCursor],
) (dto.GetPostCommentsResult, error) {
	query := builder().Select(
		"id",
		"parent_id",
		"post_id",
		"user_id",
		"content",
		"created_at",
		"updated_at",
	).Where("post_id = ?", postID)

	if crsr.Value != nil {
		switch crsr.Direction {
		case cursor.DirectionNext:
			query = query.Where("id <= ?", crsr.Value.ID)
		case cursor.DirectionPrev:
			query = query.Where("id > ?", crsr.Value.ID)
		}
	}

	query = query.OrderBy("id")
	query = query.Limit(uint64(crsr.Limit) + 1)

	sql, args, err := query.ToSql()
	if err != nil {
		return dto.GetPostCommentsResult{}, fmt.Errorf("query.ToSql: %w", err)
	}

	rows, err := db.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return dto.GetPostCommentsResult{}, fmt.Errorf("db.QueryContext: %w", err)
	}
	defer rows.Close()

	comments := make([]comment.Comment, 0, crsr.Limit+1)
	for rows.Next() {
		var comment comment.Comment

		comment, err = scanComment(rows)
		if err != nil {
			return dto.GetPostCommentsResult{}, fmt.Errorf("scanComment: %w", err)
		}

		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return dto.GetPostCommentsResult{}, fmt.Errorf("rows.Err: %w", err)
	}

	if len(comments) == 0 {
		return dto.GetPostCommentsResult{}, nil
	}

	var cursorGroup cursor.Group

	if crsr.Value != nil {
		cursorGroup.Prev = dto.GetPostCommentsCursorFromComment(comments[0]).Marshal()
	}

	if len(comments) == crsr.Limit+1 {
		cursorGroup.Next = dto.GetPostCommentsCursorFromComment(comments[len(comments)-1]).Marshal()
		comments = comments[:len(comments)-1]
	}

	return dto.GetPostCommentsResult{
		Comments:    comments,
		CursorGroup: cursorGroup,
	}, nil
}

func scanComment(row Row) (comment.Comment, error) {
	var (
		id        string
		parentID  string
		postID    string
		userID    string
		content   string
		createdAt time.Time
		updatedAt time.Time
	)

	if err := row.Scan(
		&id,
		&parentID,
		&postID,
		&userID,
		&content,
		&createdAt,
		&updatedAt,
	); err != nil {
		return comment.Comment{}, fmt.Errorf("row.Scan: %w", err)
	}

	return comment.Comment{
		ID:        comment.MustNewIDFromString(id),
		ParentID:  comment.MustNewIDFromString(id),
		PostID:    post.MustNewIDFromString(postID),
		UserID:    user.MustNewIDFromString(userID),
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
