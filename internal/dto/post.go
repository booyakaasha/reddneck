package dto

import (
	"fmt"

	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/rs/xid"
)

type GetPostsResult struct {
	Posts       []post.Post
	CursorGroup cursor.Group
}

type GetPostsCursor struct {
	ID post.ID
}

func GetPostsCursorFromPost(p post.Post) *GetPostsCursor {
	return &GetPostsCursor{
		ID: p.ID,
	}
}

func (gpc *GetPostsCursor) Marshal() string {
	return gpc.ID.String()
}

func (gpc *GetPostsCursor) Unmarshal(s string) error {
	id, err := xid.FromString(s)
	if err != nil {
		return fmt.Errorf("xid.FromString: %w", err)
	}
	gpc.ID = post.NewID(id)
	return nil
}
