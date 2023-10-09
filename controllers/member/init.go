package member

import (
	"github.com/HMUniversity/System/middleware"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
}

func (m *Controller) RegisterRouter(router fiber.Router) {
	router.Post("/join", middleware.MidAuth, join)
}
