package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"app.yourcompany.com/routes"
	"app.yourcompany.com/server"
)

// Embed a directory
//go:embed public/*
var embededPublicDir embed.FS

func main() {
	server := server.Setup()

	server.RegisterMiddlewares()

	// Register route for embedded public directory
	server.Use("/public", filesystem.New(filesystem.Config{
		Root: http.FS(embededPublicDir),
		PathPrefix: "public",
		Browse: false,
	}))

	defaultGroup := server.Group("")
	routes.RegisterWeb(defaultGroup)

	err := server.Listen()
	if err != nil {
		log.Fatal(err)
		server.Shutdown()
	}
}
