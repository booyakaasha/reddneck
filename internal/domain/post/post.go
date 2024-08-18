package post

import (
	"fmt"
	"time"

	"github.com/booyakaasha/reddneck/internal/domain/user"
	"github.com/rs/xid"
)

// ID поста.
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

// Post пост.
type Post struct {
	ID        ID
	UserID    user.ID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
