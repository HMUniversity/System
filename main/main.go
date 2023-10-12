package main

import (
	"github.com/HMUniversity/System/controllers"
	"github.com/HMUniversity/System/modules/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.Load()
	config.Get().GitHub.Init()
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		Prefork:       true,
		ServerHeader:  "HMU Web Service",
		AppName:       "HMU Core System",
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	controllers.RegisterControllers(app)
	err := app.Listen(config.Get().ListenAddr)
	if err != nil {
		panic(err)
	}
}
