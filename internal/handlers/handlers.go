package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/domain"
	"net/http"
)

// RegisterHandlers binds all the handlers to the http router
func RegisterHandlers(r *gin.Engine, blog BlogService) {
	s := server{service: blog}
	r.GET("v1/posts/:id", s.GetPostByID)
	r.GET("v1/posts", s.Posts)
	r.POST("v1/posts", s.CreatePost)
	r.DELETE("v1/posts/:id", s.DeletePost)
	r.PUT("v1/posts/:id", s.UpdatePost)
}

type BlogService interface {
	CreatePost(ctx context.Context, p *domain.Post) (domain.PostId, error)
	Post(ctx context.Context, id domain.PostId) (*domain.Post, error)
	DeletePost(ctx context.Context, id domain.PostId) error
	UpdatePost(ctx context.Context, post *domain.Post, id domain.PostId) error
	Posts(ctx context.Context) []*domain.Post
}

type server struct {
	service BlogService
}

func (s *server) GetPostByID(c *gin.Context) {
	id, err := mapPostId(c)
	if err != nil {
		c.Error(err)
		return
	}

	post, err := s.service.Post(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, post)
}

// Posts returns as the available posts.
// The API does not support pagination for sake of
// simplicity of the Storage
func (s *server) Posts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"posts": s.service.Posts(c)})
}

func (s *server) CreatePost(c *gin.Context) {
	newPost, err := mapToPost(c)
	if err != nil {
		c.Error(err)
		return
	}

	id, err := s.service.CreatePost(c.Request.Context(), newPost)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, postIdResp(id))
}

func (s *server) DeletePost(c *gin.Context) {
	id, err := mapPostId(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := s.service.DeletePost(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, msgStatusOk)
}

func (s *server) UpdatePost(c *gin.Context) {
	id, err := mapPostId(c)
	if err != nil {
		c.Error(err)
		return
	}

	post, err := mapToPost(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := s.service.UpdatePost(c.Request.Context(), post, id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, postIdResp(id))
}
