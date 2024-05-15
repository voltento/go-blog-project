package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/httperr"
	"net/http"
)

const (
	invalidStatusCode = 0
)

// HttpErrHandlerMiddleware handles http errors to extract status code
func HttpErrHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := invalidStatusCode
		var httpErr error
		for _, err := range c.Errors {
			code := httperr.HTTPStatusCode(err, http.StatusInternalServerError)
			if statusCode == invalidStatusCode {
				statusCode = code
				httpErr = err
			}
		}

		if statusCode != 0 {
			c.JSON(statusCode, gin.H{"error": httpErr.Error()})
		}
	}
}
