package server

import (
	"go_padawan/src/server/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	port   string
	server *fiber.App
}

func NewServer() server {
	return server{
		port:   ":5000",
		server: fiber.New(),
	}
}

func (s *server) Run() {
	router := routes.ConfigRoutes(s.server)

	router.Listen(s.port)
	log.Print("Server running on port ", s.port)

}
