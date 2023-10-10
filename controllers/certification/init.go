package certification

import "github.com/gofiber/fiber/v2"

type Controller struct {
}

func (m *Controller) RegisterRouter(router fiber.Router) {
	router.Post("/generate", generate)
}
