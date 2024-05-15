package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/blog"
	"github.com/voltento/go-blog-project/internal/handlers"
	"github.com/voltento/go-blog-project/internal/middlewares"
	"github.com/voltento/go-blog-project/internal/storage"
)

func main() {
	port := flag.String("port", "8080", "Port for the API handlers")
	flag.Parse()

	r := gin.New()
	middlewares.Setup(r)

	s := storage.NewStorage()

	b := blog.NewBlog(s)
	handlers.RegisterPostHandlers(r, b)

	_ = r.Run(":" + *port)
}
