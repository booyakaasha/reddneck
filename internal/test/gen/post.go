package gen

import (
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/domain/user"
	"github.com/rs/xid"
)

func NewPost() post.Post {
	return post.Post{
		ID:        post.NewID(xid.New()),
		UserID:    user.NewID(xid.New()),
		Title:     "title",
		Content:   "content",
		CreatedAt: Now(),
		UpdatedAt: Now(),
	}
}
