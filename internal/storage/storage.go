package storage

import (
	"fmt"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/internal/httperr"
	"net/http"
	"sync"
	"sync/atomic"
)

type Storage struct {
	postsMtx sync.RWMutex
	posts    map[domain.PostId]*domain.Post

	seqId *int64
}

func (s *Storage) UpdatePost(post *domain.Post, id domain.PostId) error {
	post.ID = id

	s.postsMtx.Lock()
	defer s.postsMtx.Unlock()

	if _, exists := s.posts[id]; !exists {
		err := fmt.Errorf("blog not found. id: %v", id)
		return httperr.WrapWithHttpCode(err, http.StatusNotFound)
	}

	s.posts[id] = post
	return nil
}

func (s *Storage) DeletePost(id domain.PostId) error {
	s.postsMtx.Lock()
	defer s.postsMtx.Unlock()

	delete(s.posts, id)
	return nil
}

func (s *Storage) Post(id domain.PostId) (*domain.Post, error) {
	s.postsMtx.RLock()
	defer s.postsMtx.RUnlock()

	if p, isOk := s.posts[id]; isOk {
		return p, nil
	}

	err := fmt.Errorf("blog not found. id: %v", id)
	return nil, httperr.WrapWithHttpCode(err, http.StatusNotFound)
}

func (s *Storage) CreatePost(post *domain.Post) (domain.PostId, error) {
	nextPostId := s.nextAvailableId()
	post.ID = nextPostId

	s.postsMtx.Lock()
	defer s.postsMtx.Unlock()

	s.posts[nextPostId] = post
	return nextPostId, nil
}

// nextAvailableId returns next post id which is guarantied to be not used yet
func (s *Storage) nextAvailableId() domain.PostId {
	s.postsMtx.RLock()
	defer s.postsMtx.RUnlock()

	for {
		// Acquire next available seq Id
		seqId := atomic.AddInt64(s.seqId, 1)
		postId := domain.PostId(seqId - 1)

		if _, busy := s.posts[postId]; !busy {
			return postId
		}
	}
}

func NewStorage() *Storage {
	var startId int64 = 1
	return &Storage{
		seqId: &startId,
		posts: map[domain.PostId]*domain.Post{},
	}
}
