package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/blog"
	"github.com/voltento/go-blog-project/internal/handlers"
	"github.com/voltento/go-blog-project/internal/middlewares"
	"github.com/voltento/go-blog-project/internal/migration"
	"github.com/voltento/go-blog-project/internal/storage"
	"golang.org/x/exp/slog"
)

func main() {
	port := flag.String("port", "8080", "Port for the API handlers")
	migrationFile := flag.String("migration", "./resourses/blog_data.json", "Migration file")
	flag.Parse()

	s := storage.NewStorage()

	if len(*migrationFile) > 1 {
		slog.Info("migration started", "migration file", *migrationFile)
		m := migration.Migration{}
		err := m.Apply(context.Background(), *migrationFile, s)
		if err != nil {
			slog.Error("can not apply migration.", "error", err)
			return
		}

		slog.Info("migration applied")
	}

	r := gin.New()
	middlewares.Setup(r)

	b := blog.NewBlog(s)
	handlers.RegisterHandlers(r, b)

	_ = r.Run(":" + *port)
}
