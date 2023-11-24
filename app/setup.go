package app

import (
	"fmt"
	"log"
	"os"

	"github.com/SoNim-LSCM/maxbot_oms/config"
	errorHandler "github.com/SoNim-LSCM/maxbot_oms/errors"
	"github.com/SoNim-LSCM/maxbot_oms/mqtt"

	// "github.com/SoNim-LSCM/maxbot_oms/mqtt"

	"github.com/SoNim-LSCM/maxbot_oms/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() {
	// load env
	err := config.LoadENV()
	errorHandler.CheckError(err, "load env")

	// set output logs
	var f *os.File
	go config.SetupLogCron(f)

	defer f.Close()

	log.SetOutput(f)
	log.Println("SYSTEM RESTARTED")

	// start database
	// go database.StartMySql()
	// errorHandler.CheckError(err, "Start MySql")

	fmt.Println("1")
	// start mqtt server
	go mqtt.MqttSetup()
	errorHandler.CheckError(err, "start MQTT")

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	// config.AddSwaggerRoutes(app)

	// setup websocket
	// go websocket.SetupWebsocket()

	// get the port and start
	port := os.Getenv("API_PORT")
	app.Listen(":" + port)

	log.Println("FINISH SYSTEM CONFIG")
}
