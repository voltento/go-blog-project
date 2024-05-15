package blog

import (
	"context"
	"github.com/voltento/go-blog-project/internal/domain"
)

type Storage interface {
	Post(ctx context.Context, id domain.PostId) (*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) (domain.PostId, error)
	DeletePost(ctx context.Context, id domain.PostId) error
	UpdatePost(ctx context.Context, post *domain.Post, id domain.PostId) error
}

type Blog struct {
	storage Storage
}

func NewBlog(s Storage) *Blog {
	return &Blog{storage: s}
}

func (b *Blog) CreatePost(ctx context.Context, p *domain.Post) (domain.PostId, error) {
	return b.storage.CreatePost(ctx, p)
}

func (b *Blog) Post(ctx context.Context, id domain.PostId) (*domain.Post, error) {
	return b.storage.Post(ctx, id)
}

func (b *Blog) DeletePost(ctx context.Context, id domain.PostId) error {
	return b.storage.DeletePost(ctx, id)
}

func (b *Blog) UpdatePost(ctx context.Context, post *domain.Post, id domain.PostId) error {
	return b.storage.UpdatePost(ctx, post, id)
}
