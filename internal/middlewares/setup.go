package middlewares

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	r.Use(LoggerMiddleware())
	r.Use(HttpErrHandlerMiddleware())
	r.Use(gin.Recovery())
}
