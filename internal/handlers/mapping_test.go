package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/internal/httperr"
)

func TestMapPostId(t *testing.T) {
	tests := []struct {
		name         string
		param        string
		expectedID   domain.PostId
		expectedErr  bool
		expectedCode int
	}{
		{"Valid ID", "123", domain.PostId(123), false, 0},
		{"Invalid ID", "abc", 0, true, http.StatusBadRequest},
		{"Empty ID", "", 0, true, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Params = gin.Params{gin.Param{Key: "id", Value: tt.param}}

			id, err := mapPostId(c)

			if tt.expectedErr {
				assert.Error(t, err)
				httpStatus := httperr.HTTPStatusCode(err, -1)
				assert.Equal(t, tt.expectedCode, httpStatus)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}
		})
	}
}

func TestMapToPost(t *testing.T) {
	tests := []struct {
		name         string
		jsonPayload  string
		expectedPost *domain.Post
		expectedErr  bool
		expectedCode int
	}{
		{
			"Valid JSON Payload",
			`{"title":"New Post","content":"New Content","author":"New Author"}`,
			&domain.Post{Title: "New Post", Content: "New Content", Author: "New Author"},
			false,
			0,
		},
		{
			"Invalid JSON Payload",
			`{"title":"New Post","content":123,"author":"New Author"}`,
			nil,
			true,
			http.StatusBadRequest,
		},
		{
			"Missing author",
			`{"title":"New Post","content":"New Content"}`,
			nil,
			true,
			http.StatusBadRequest,
		},
		{
			"Missing title",
			`{"content":"New Content","author":"New Author"}`,
			nil,
			true,
			http.StatusBadRequest,
		},
		{
			"Missing content",
			`{"title":"New Post", "author":"New Author"}`,
			nil,
			true,
			http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = httptest.NewRequest("POST", "/", strings.NewReader(tt.jsonPayload))
			c.Request.Header.Set("Content-Type", "application/json")

			post, err := mapToPost(c)

			if tt.expectedErr {
				assert.Error(t, err)
				httpStatus := httperr.HTTPStatusCode(err, -1)
				assert.Equal(t, tt.expectedCode, httpStatus)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPost, post)
			}
		})
	}
}
