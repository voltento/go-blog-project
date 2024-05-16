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

	"github.com/gavv/httpexpect/v2"
)

type HandlersTestSuite struct {
	suite.Suite
	server   *httptest.Server
	mockBlog *mocks.BlogService
	router   *gin.Engine

	expect *httpexpect.Expect
}

func (s *HandlersTestSuite) SetupTest() {
	s.mockBlog = new(mocks.BlogService)

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()

	middlewares.Setup(s.router)
	RegisterHandlers(s.router, s.mockBlog)
	s.server = httptest.NewServer(s.router)

	s.expect = httpexpect.Default(s.T(), s.server.URL)
}

func (s *HandlersTestSuite) TestGetPostByID() {
	s.mockBlog.On("Post", mock.Anything, domain.PostId(1)).Return(&domain.Post{ID: 1, Title: "Test Title", Content: "Test Content", Author: "Test Author"}, nil)

	s.expect.GET("/posts/1").
		Expect().
		Status(http.StatusOK).
		Body().IsEqual("{\"ID\":1,\"Title\":\"Test Title\",\"Content\":\"Test Content\",\"Author\":\"Test Author\"}")

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestGetPostByID_WrongPostIdFormat() {
	s.expect.GET("GET", "/posts/wrongId").
		Expect().
		Status(http.StatusBadRequest)
}

func (s *HandlersTestSuite) TestCreatePost() {
	newPost := &domain.Post{Title: "New Post", Content: "New Content", Author: "New Author"}
	s.mockBlog.On("CreatePost", mock.Anything, mock.Anything).Return(domain.PostId(1), nil).Run(func(args mock.Arguments) {
		post := args.Get(1).(*domain.Post)
		s.Equal(newPost.Title, post.Title)
		s.Equal(newPost.Content, post.Content)
		s.Equal(newPost.Author, post.Author)
	})

	s.expect.POST("/posts").
		WithJSON(newPost).
		Expect().
		Status(http.StatusCreated).
		Body().IsEqual("{\"postId\":1}")

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestCreatePost_WrongPostFormat() {
	s.expect.POST("/posts").
		WithJSON(map[string]string{"content": "New Content", "author": "New Author"}).
		Expect().
		Status(http.StatusBadRequest)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestCreatePost_ServiceCreateRequestFailed() {
	err := httperr.WrapWithHttpCode(errors.New(""), http.StatusConflict)
	s.mockBlog.On("CreatePost", mock.Anything, mock.Anything).Return(domain.PostId(1), err)

	s.expect.POST("/posts").
		WithJSON(map[string]string{"title": "New Post", "content": "New Content", "author": "New Author"}).
		Expect().
		Status(http.StatusConflict)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost() {
	s.mockBlog.On("DeletePost", mock.Anything, domain.PostId(1)).Return(nil)

	s.expect.DELETE("/posts/1").Expect().
		Status(http.StatusNoContent)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost_WrongPostIdFormat() {
	s.expect.DELETE("/posts/wrongId").Expect().
		Status(http.StatusBadRequest)

	s.mockBlog.AssertExpectations(s.T())
}

func (s *HandlersTestSuite) TestDeletePost_serviceReturnsError() {
	err := httperr.WrapWithHttpCode(errors.New(""), http.StatusConflict)
	s.mockBlog.On("DeletePost", mock.Anything, domain.PostId(1)).Return(err)

	s.expect.DELETE("/posts/1").Expect().
		Status(http.StatusConflict)

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

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
