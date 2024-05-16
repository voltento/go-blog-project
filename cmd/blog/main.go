package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/voltento/go-blog-project/internal/blog"
	"github.com/voltento/go-blog-project/internal/handlers"
	"github.com/voltento/go-blog-project/internal/middlewares"
	"github.com/voltento/go-blog-project/internal/migration"
	"github.com/voltento/go-blog-project/internal/storage"
	"golang.org/x/exp/slog"
	"log"
	"os"
)

func run() error {
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
			return err
		}

		slog.Info("migration applied")
	}

	r := gin.New()
	middlewares.Setup(r)

	b := blog.NewBlog(s)
	handlers.RegisterHandlers(r, b)

	return r.Run(":" + *port)
}

//	@Title			Blog API
//	@Version		1.0
//	@Description	This is a Simple Blog Server
//
// @host	localhost:8080
func main() {
	if err := run(); err != nil {
		log.Fatal("Failed to start", "error", err)
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
