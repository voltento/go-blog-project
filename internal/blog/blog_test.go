package blog

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/mocks"
)

type BlogTestSuite struct {
	suite.Suite
	mockStorage *mocks.Storage
	blog        *Blog
	ctx         context.Context
	postId      domain.PostId
}

func (s *BlogTestSuite) SetupTest() {
	s.mockStorage = new(mocks.Storage)
	s.blog = NewBlog(s.mockStorage)
	s.ctx = context.Background()
	s.postId = domain.PostId(1)
}

func (s *BlogTestSuite) TestCreatePost() {
	newPost := &domain.Post{Title: "New Post", Content: "New Content", Author: "New Author"}

	s.mockStorage.On("CreatePost", s.ctx, newPost).Return(s.postId, nil)

	id, err := s.blog.CreatePost(s.ctx, newPost)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.postId, id)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestCreatePost_Error() {
	newPost := &domain.Post{Title: "New Post", Content: "New Content", Author: "New Author"}
	s.mockStorage.On("CreatePost", s.ctx, newPost).Return(domain.PostId(0), errors.New("error"))

	id, err := s.blog.CreatePost(s.ctx, newPost)

	assert.Error(s.T(), err)
	assert.Equal(s.T(), domain.PostId(0), id)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestPost() {
	post := &domain.Post{ID: 1, Title: "Test Post", Content: "Test Content", Author: "Test Author"}
	postId := domain.PostId(1)

	s.mockStorage.On("Post", s.ctx, postId).Return(post, nil)

	result, err := s.blog.Post(s.ctx, postId)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), post, result)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestPost_Error() {
	postId := domain.PostId(1)
	s.mockStorage.On("Post", s.ctx, postId).Return(nil, errors.New("error"))

	result, err := s.blog.Post(s.ctx, postId)

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestDeletePost() {
	postId := domain.PostId(1)

	s.mockStorage.On("DeletePost", s.ctx, postId).Return(nil)

	err := s.blog.DeletePost(s.ctx, postId)

	assert.NoError(s.T(), err)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestDeletePost_Error() {
	postId := domain.PostId(1)
	s.mockStorage.On("DeletePost", s.ctx, postId).Return(errors.New("error"))

	err := s.blog.DeletePost(s.ctx, postId)

	assert.Error(s.T(), err)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestUpdatePost() {
	updatedPost := &domain.Post{ID: 1, Title: "Updated Post", Content: "Updated Content", Author: "Updated Author"}
	postId := domain.PostId(1)

	s.mockStorage.On("UpdatePost", s.ctx, updatedPost, postId).Return(nil)

	err := s.blog.UpdatePost(s.ctx, updatedPost, postId)

	assert.NoError(s.T(), err)
	s.mockStorage.AssertExpectations(s.T())
}

func (s *BlogTestSuite) TestUpdatePost_Error() {
	updatedPost := &domain.Post{ID: 1, Title: "Updated Post", Content: "Updated Content", Author: "Updated Author"}

	s.mockStorage.On("UpdatePost", s.ctx, updatedPost, s.postId).Return(errors.New("error"))

	err := s.blog.UpdatePost(s.ctx, updatedPost, s.postId)

	assert.Error(s.T(), err)
	s.mockStorage.AssertExpectations(s.T())
}

func TestBlogTestSuite(t *testing.T) {
	suite.Run(t, new(BlogTestSuite))
}
