package repository

import (
	"context"

	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/comment"
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/dto"
)

type Comment interface {
	CreateComment(ctx context.Context, comment comment.Comment) error
	GetCommentByID(ctx context.Context, id comment.ID) (comment.Comment, error)
	GetPostComments(
		ctx context.Context,
		postID post.ID,
		crsr cursor.Cursor[*dto.GetPostCommentsCursor],
	) (comment.Comment, error)
}
