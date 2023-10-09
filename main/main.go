package main

import (
	"github.com/HMUniversity/System/controllers"
	"github.com/HMUniversity/System/modules/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	config.Get().GitHub.Init()
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "HMU Web Service",
		AppName:       "HMU Core System",
	})
	controllers.RegisterControllers(app)
	err := app.Listen(config.Get().ListenAddr)
	if err != nil {
		panic(err)
	}
}
