package router

import (
	"github.com/SoNim-LSCM/maxbot_oms/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// setup the oms group
	maxbot := app.Group("/maxbot")

	// Health Check
	maxbot.Get("/addOrder", handlers.HandleAddOrder)

}
