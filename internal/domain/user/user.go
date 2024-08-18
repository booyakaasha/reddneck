package user

import (
	"fmt"

	"github.com/rs/xid"
)

// ID пользователя.
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
