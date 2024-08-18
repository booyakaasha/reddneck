package usecase_test

import (
	"context"
	"testing"

	"github.com/booyakaasha/reddneck/internal/test"
	"github.com/booyakaasha/reddneck/internal/test/gen"
	"github.com/booyakaasha/reddneck/internal/usecase"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostUsecase(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	pu := usecase.NewPostUsecase(
		test.NewDB(t),
	)

	post := gen.NewPost()
	require.NoError(t, pu.CreatePost(ctx, post))
	require.NoError(t, pu.CreatePost(ctx, post), "идемпотентный вызов")

	actual, err := pu.GetPostByID(ctx, post.ID)
	require.NoError(t, err)

	assert.Empty(
		t,
		cmp.Diff(
			post,
			actual,
		),
	)
}
