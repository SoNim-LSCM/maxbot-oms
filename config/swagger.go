package config

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/SoNim-LSCM/maxbot_oms/docs"
)

func AddSwaggerRoutes(app *fiber.App) {
	// setup swagger
	app.Get("/oms/swagger/*", fiberSwagger.WrapHandler)
}
