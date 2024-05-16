package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/voltento/go-blog-project/internal/domain"
	"os"
)

type Storage interface {
	CreatePost(ctx context.Context, post *domain.Post) (domain.PostId, error)
}

type Migration struct {
}

func (m *Migration) Apply(ctx context.Context, filePath string, s Storage) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	blogData := struct {
		Posts []post `json:"Posts"`
	}{}
	if err := json.Unmarshal(data, &blogData); err != nil {
		return err
	}

	for _, p := range blogData.Posts {
		_, err := s.CreatePost(ctx, &domain.Post{
			ID:      domain.PostId(p.ID),
			Title:   p.Title,
			Content: p.Content,
			Author:  p.Author,
		})

		if err != nil {
			return fmt.Errorf("can not apply migration from file '%s'. error: %w", filePath, err)
		}
	}

	return nil
}

type post struct {
	ID      int    `json:"ID,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
