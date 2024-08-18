package usecase

import (
	"context"
	"fmt"

	"github.com/booyakaasha/reddneck/internal/domain/post"
	"github.com/booyakaasha/reddneck/internal/repository"
)

type PostUsecase struct {
	repository repository.Post
}

func NewPostUsecase(
	repository repository.Post,
) PostUsecase {
	return PostUsecase{
		repository: repository,
	}
}

func (pu PostUsecase) CreatePost(ctx context.Context, post post.Post) error {
	if err := pu.repository.CreatePost(ctx, post); err != nil {
		return fmt.Errorf("repository.CreatePost: %w", err)
	}

	return nil
}

func (pu PostUsecase) GetPostByID(ctx context.Context, id post.ID) (post.Post, error) {
	p, err := pu.repository.GetPostByID(ctx, id)
	if err != nil {
		return post.Post{}, fmt.Errorf("repository.GetPostByID: %w", err)
	}

	return p, nil
}
