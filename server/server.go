package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

type Server struct {
	*fiber.App
}

func Setup() *Server {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Server{fiber.New()}
}

func (s *Server) Listen() error {
	listenAddr := HttpListenAddress()

	fmt.Printf("Application running in %s at http://%s\n", Env(), listenAddr)
	
	return s.App.Listen(listenAddr)
}

func (s *Server) Shutdown() {
	err := s.App.Shutdown()
	if err != nil {
		os.Exit(1)
	}
}

func(s *Server) RegisterMiddlewares() {
	// Register default panic recovery middleware
	s.Use(recover.New())
}

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func Env() string {
	return os.Getenv("APP_ENV")
}

func HttpListenAddress() string {
	return os.Getenv("HTTP_LISTEN_ADDR")
}