package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/internal/httperr"
	"github.com/voltento/go-blog-project/internal/middlewares"
	"github.com/voltento/go-blog-project/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandlersTestSuite struct {
	suite.Suite
	mockBlog *mocks.BlogService
	router   *gin.Engine
}

func (s *HandlersTestSuite) SetupTest() {
	s.mockBlog = new(mocks.BlogService)
	s.router = setupRouter(s.mockBlog)
}

func (s *HandlersTestSuite) TestGetPostByID() {
	s.mockBlog.On("Post", mock.Anything, domain.PostId(1)).Return(&domain.Post{ID: 1, Title: "Test Title", Content: "Test Content", Author: "Test Author"}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/1", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Contains(s.T(), w.Body.String(), "Test Title")
	assert.Contains(s.T(), w.Body.String(), "Test Content")
	assert.Contains(s.T(), w.Body.String(), "Test Author")

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestGetPostByID_WrongPostIdFormat() {
	mockBlog := new(mocks.BlogService)
	router := setupRouter(mockBlog)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/wrongId", nil)
	router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestCreatePost() {
	newPost := &domain.Post{Title: "New Post", Content: "New Content", Author: "New Author"}
	s.mockBlog.On("CreatePost", mock.Anything, mock.Anything).Return(domain.PostId(1), nil).Run(func(args mock.Arguments) {
		post := args.Get(1).(*domain.Post)
		assert.Equal(s.T(), newPost.Title, post.Title)
		assert.Equal(s.T(), newPost.Content, post.Content)
		assert.Equal(s.T(), newPost.Author, post.Author)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(`{"title":"New Post","content":"New Content","author":"New Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusCreated, w.Code)
	assert.Contains(s.T(), w.Body.String(), `"postId":1`)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestCreatePost_WrongPostFormat() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(`{"content":"New Content","author":"New Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestCreatePost_ServiceCreateRequestFailed() {
	err := httperr.WrapWithHttpCode(errors.New(""), http.StatusConflict)
	s.mockBlog.On("CreatePost", mock.Anything, mock.Anything).Return(domain.PostId(1), err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(`{"title":"New Post","content":"New Content","author":"New Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusConflict, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost() {
	s.mockBlog.On("DeletePost", mock.Anything, domain.PostId(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/posts/1", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusNoContent, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost_WrongPostIdFormat() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/posts/wrongId", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost_serviceReturnsError() {
	err := httperr.WrapWithHttpCode(errors.New(""), http.StatusConflict)
	s.mockBlog.On("DeletePost", mock.Anything, domain.PostId(1)).Return(err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/posts/1", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusConflict, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestUpdatePost() {
	s.mockBlog.On("UpdatePost", mock.Anything, mock.Anything, domain.PostId(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/posts/1", strings.NewReader(`{"title":"Updated Title","content":"Updated Content","author":"Updated Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Contains(s.T(), w.Body.String(), `"postId":1`)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestUpdatePost_WrongPostIdFormat() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/posts/wrongId", strings.NewReader(`{"title":"Updated Title","content":"Updated Content","author":"Updated Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestUpdatePost_serviceReturnsError() {
	err := httperr.WrapWithHttpCode(errors.New(""), http.StatusForbidden)
	s.mockBlog.On("UpdatePost", mock.Anything, mock.Anything, domain.PostId(1)).Return(err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/posts/1", strings.NewReader(`{"title":"Updated Title","content":"Updated Content","author":"Updated Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusForbidden, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestUpdatePost_wrongPostFormat() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/posts/1", strings.NewReader(`faulty format`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusBadRequest, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestGetPostByID_NotFound() {
	s.mockBlog.On("Post", mock.Anything, domain.PostId(1)).Return(&domain.Post{}, httperr.WrapWithHttpCode(errors.New("post not found"), http.StatusNotFound))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts/1", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusNotFound, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestUpdatePost_NotFound() {
	s.mockBlog.On("UpdatePost", mock.Anything, mock.Anything, domain.PostId(1)).Return(httperr.WrapWithHttpCode(errors.New("post not found"), http.StatusNotFound))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/posts/1", strings.NewReader(`{"title":"Updated Title","content":"Updated Content","author":"Updated Author"}`))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusNotFound, w.Code)

	s.mockBlog.AssertExpectations(s.T())
}

func setupRouter(blogService *mocks.BlogService) *gin.Engine {
	r := gin.Default()
	middlewares.Setup(r)
	RegisterHandlers(r, blogService)
	return r
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
