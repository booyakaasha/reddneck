package comment

import (
	"fmt"
	"time"

	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/domain/user"
	"github.com/rs/xid"
)

// ID id комментария.
type ID struct {
	xid.ID
}

func NewID(id xid.ID) ID {
	return ID{
		ID: id,
	}
}

func MustNewIDFromString(s string) ID {
	id, err := xid.FromString(s)
	if err != nil {
		panic(fmt.Sprintf("xid.FromString: %v", err))
	}
	return NewID(id)
}

// Comment комментарий к посту.
type Comment struct {
	ID        ID
	ParentID  ID
	PostID    post.ID
	UserID    user.ID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
