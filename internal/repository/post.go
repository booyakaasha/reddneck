package repository

import (
	"context"

	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/dto"
)

type Post interface {
	CreatePost(ctx context.Context, post post.Post) error
	GetPostByID(ctx context.Context, id post.ID) (post.Post, error)
	GetPosts(ctx context.Context, crsr cursor.Cursor[*dto.GetPostsCursor]) (dto.GetPostsResult, error)
}
