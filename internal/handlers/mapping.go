package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/domain"
	"github.com/voltento/go-blog-project/internal/httperr"
	"net/http"
	"strconv"
)

var msgStatusOk = gin.H{"status": "ok"}

func postIdResp(id domain.PostId) gin.H {
	return gin.H{"postId": id}
}

type PostDTO struct {
	ID      domain.PostId `json:"id"`
	Title   string        `json:"title" binding:"required"`
	Content string        `json:"content" binding:"required"`
	Author  string        `json:"author" binding:"required"`
}

func mapPostId(c *gin.Context) (domain.PostId, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = fmt.Errorf("can not parse post id '%s'. error: %w", idStr, err)
		return 0, httperr.WrapWithHttpCode(err, http.StatusBadRequest)
	}

	return domain.PostId(id), nil
}

func mapToPost(c *gin.Context) (*domain.Post, error) {
	var newPost PostDTO
	if err := c.BindJSON(&newPost); err != nil {
		err = httperr.WrapWithHttpCode(err, http.StatusBadRequest)
		return nil, err
	}

	return &domain.Post{
		ID:      newPost.ID,
		Title:   newPost.Title,
		Content: newPost.Content,
		Author:  newPost.Author,
	}, nil
}
