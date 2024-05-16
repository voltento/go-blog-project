package migration

import (
	"context"
	"github.com/voltento/go-blog-project/internal/domain"
)

type Storage interface {
	CreatePost(ctx context.Context, post *domain.Post) (domain.PostId, error)
}
