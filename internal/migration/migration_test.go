package migration

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/mocks"
	"os"
	"testing"
)

func TestMigration_Apply(t *testing.T) {
	ctx := context.Background()

	t.Run("successful migration", func(t *testing.T) {
		mockStorage := new(mocks.Storage)
		migration := &Migration{}

		data := struct {
			Posts []post `json:"Posts"`
		}{Posts: []post{
			{ID: 1, Title: "Title 1", Content: "Content 1", Author: "Author 1"},
			{ID: 2, Title: "Title 2", Content: "Content 2", Author: "Author 2"},
		}}

		filename, clean := tempFileFromData(t, data)
		defer clean()

		mockStorage.On("CreatePost", ctx, &domain.Post{
			ID:      domain.PostId(1),
			Title:   "Title 1",
			Content: "Content 1",
			Author:  "Author 1",
		}).Return(domain.PostId(1), nil).Once()

		mockStorage.On("CreatePost", ctx, &domain.Post{
			ID:      domain.PostId(2),
			Title:   "Title 2",
			Content: "Content 2",
			Author:  "Author 2",
		}).Return(domain.PostId(2), nil).Once()

		err := migration.Apply(ctx, filename, mockStorage)
		assert.NoError(t, err)
		mockStorage.AssertExpectations(t)
	})

	t.Run("file read error", func(t *testing.T) {
		mockStorage := new(mocks.Storage)
		migration := &Migration{}
		filePath := "nonexistent_file.json"

		err := migration.Apply(ctx, filePath, mockStorage)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no such file or directory")
	})

	t.Run("json unmarshal error", func(t *testing.T) {
		mockStorage := new(mocks.Storage)
		migration := &Migration{}

		tempFile, err := os.CreateTemp("", "invalid.json")
		assert.NoError(t, err)
		defer os.Remove(tempFile.Name())

		_, err = tempFile.Write([]byte("invalid json"))
		assert.NoError(t, err)
		assert.NoError(t, tempFile.Close())

		err = migration.Apply(ctx, tempFile.Name(), mockStorage)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})

	t.Run("create post error", func(t *testing.T) {
		mockStorage := new(mocks.Storage)
		migration := &Migration{}

		data := struct {
			Posts []post `json:"Posts"`
		}{Posts: []post{
			{ID: 1, Title: "Title 1", Content: "Content 1", Author: "Author 1"},
			{ID: 2, Title: "Title 2", Content: "Content 2", Author: "Author 2"},
		}}

		filename, clean := tempFileFromData(t, data)
		defer clean()

		mockStorage.On("CreatePost", ctx, &domain.Post{
			ID:      domain.PostId(1),
			Title:   "Title 1",
			Content: "Content 1",
			Author:  "Author 1",
		}).Return(domain.PostId(0), errors.New("failed to create post")).Once()

		err := migration.Apply(ctx, filename, mockStorage)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "can not apply migration from file")
		mockStorage.AssertExpectations(t)
	})
}

type cleanF func()

func tempFileFromData(t *testing.T, vs any) (string, cleanF) {
	data, err := json.Marshal(vs)
	assert.NoError(t, err)

	tempFile, err := os.CreateTemp("", "Posts.json")
	assert.NoError(t, err)

	_, err = tempFile.Write(data)
	assert.NoError(t, err)
	assert.NoError(t, tempFile.Close())

	return tempFile.Name(), func() { _ = os.Remove(tempFile.Name()) }
}
