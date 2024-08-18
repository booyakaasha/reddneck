package dto

import (
	"fmt"

	"github.com/booyakaasha/reddneck/internal/cursor"
	"github.com/booyakaasha/reddneck/internal/domain/comment"
	"github.com/rs/xid"
)

type GetPostCommentsResult struct {
	Comments    []comment.Comment
	CursorGroup cursor.Group
}

type GetPostCommentsCursor struct {
	ID comment.ID
}

func GetPostCommentsCursorFromComment(c comment.Comment) *GetPostCommentsCursor {
	return &GetPostCommentsCursor{
		ID: c.ID,
	}
}

func (gcc *GetPostCommentsCursor) Marshal() string {
	return gcc.ID.String()
}

func (gcc *GetPostCommentsCursor) Unmarshal(s string) error {
	id, err := xid.FromString(s)
	if err != nil {
		return fmt.Errorf("xid.FromString: %w", err)
	}
	gcc.ID = comment.NewID(id)
	return nil
}

// Cursor для комментариев.
type CommentCursor struct {
	Direction cursor.Direction
	Value     GetPostCommentsCursor
	Limit     int
}

// NewGroup создает новую группу курсоров для комментариев.
func NewGroup(next, prev string) cursor.Group {
	return cursor.Group{
		Next: next,
		Prev: prev,
	}
}
