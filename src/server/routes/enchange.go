package routes

import (
	"go_padawan/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(route *fiber.App) *fiber.App {
	route.Get("exchange/:amount/:from/:to/:rate", controllers.ExchangeValue)
	return route
}
