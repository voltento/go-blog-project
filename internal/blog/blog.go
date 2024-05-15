package blog

import "github.com/voltento/go-blog-project/internal/domain"

type Storage interface {
	Post(id domain.PostId) (*domain.Post, error)
	CreatePost(post *domain.Post) (domain.PostId, error)
	DeletePost(id domain.PostId) error
	UpdatePost(post *domain.Post, id domain.PostId) error
}

type Blog struct {
	storage Storage
}

func NewBlog(s Storage) *Blog {
	return &Blog{storage: s}
}

func (b *Blog) CreatePost(p *domain.Post) (domain.PostId, error) {
	return b.storage.CreatePost(p)
}

func (b *Blog) Post(id domain.PostId) (*domain.Post, error) {
	return b.storage.Post(id)
}

func (b *Blog) DeletePost(id domain.PostId) error {
	return b.storage.DeletePost(id)
}

func (b *Blog) UpdatePost(post *domain.Post, id domain.PostId) error {
	return b.storage.UpdatePost(post, id)
}
