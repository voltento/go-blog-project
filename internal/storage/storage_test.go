package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/internal/httperr"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	s := NewStorage()
	s.posts = make(map[domain.PostId]*domain.Post)
	s.posts[1] = &domain.Post{ID: 1, Title: "Test Post", Content: "Content", Author: "Author"}

	t.Run("Post exists", func(t *testing.T) {
		post, err := s.Post(1)
		assert.NoError(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, domain.PostId(1), post.ID)
		assert.Equal(t, "Test Post", post.Title)
	})

	t.Run("Post does not exist", func(t *testing.T) {
		post, err := s.Post(2)
		assert.Nil(t, post)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, httperr.HTTPStatusCode(err, -1))
	})
}

func TestCreatePost(t *testing.T) {
	s := NewStorage()
	s.posts = make(map[domain.PostId]*domain.Post)
	post := &domain.Post{Title: "New Post", Content: "New Content", Author: "New Author", ID: 10}

	t.Run("Create a new post", func(t *testing.T) {
		id, err := s.CreatePost(post)

		assert.NoError(t, err)
		assert.NotZero(t, id)

		createdPost, err := s.Post(id)
		assert.NoError(t, err)
		assert.Equal(t, "New Post", createdPost.Title)
		assert.Equal(t, "New Content", createdPost.Content)
		assert.Equal(t, "New Author", createdPost.Author)
		assert.Equal(t, id, createdPost.ID)
	})
}

func TestNextAvailableId(t *testing.T) {
	s := NewStorage()
	s.posts = make(map[domain.PostId]*domain.Post)
	post := &domain.Post{}

	t.Run("Next available ID", func(t *testing.T) {
		id := s.nextAvailableId()

		// Simulate creating a post to consume an ID
		s.posts[id] = &domain.Post{ID: id}
		assert.Equal(t, s.nextAvailableId(), id+1)
	})

	t.Run("Next available ID skips busy values", func(t *testing.T) {
		id := s.nextAvailableId()

		// Simulate creating a post to consume an ID
		s.posts[id] = post
		s.posts[id+1] = post

		assert.Equal(t, s.nextAvailableId(), id+2, "id+1 should be skipped")
	})
}

func TestNewStorage(t *testing.T) {
	s := NewStorage()

	assert.NotNil(t, s)
	assert.NotNil(t, s.seqId)
	assert.Empty(t, s.posts)

	// Initialize posts map
	s.posts = make(map[domain.PostId]*domain.Post)
	assert.NotNil(t, s.posts)
}

func TestStorage_DeletePost(t *testing.T) {
	s := NewStorage()
	s.posts = make(map[domain.PostId]*domain.Post)

	t.Run("Delete existing post", func(t *testing.T) {
		post := &domain.Post{}
		id, err := s.CreatePost(post)
		assert.NoError(t, err)

		assert.NoError(t, s.DeletePost(id))
	})

	t.Run("Delete not existing post", func(t *testing.T) {
		id := s.nextAvailableId()
		assert.NoError(t, s.DeletePost(id))
	})
}

func TestStorage_UpdatePost(t *testing.T) {
	s := NewStorage()
	s.posts = make(map[domain.PostId]*domain.Post)

	t.Run("Update existing post", func(t *testing.T) {
		id, err := s.CreatePost(&domain.Post{})
		assert.NoError(t, err)

		post := &domain.Post{
			Title:   "title",
			Content: "content",
			Author:  "author",
		}

		assert.NoError(t, s.UpdatePost(post, id))

		updatedPost, err := s.Post(id)
		assert.NoError(t, err)

		assert.EqualValues(t, post, updatedPost)
	})

	t.Run("Update not existing post", func(t *testing.T) {
		id := s.nextAvailableId()
		post := &domain.Post{}

		err := s.UpdatePost(post, id)
		assert.Error(t, err)

		assert.Equal(t, http.StatusNotFound, httperr.HTTPStatusCode(err, -1))
	})
}
