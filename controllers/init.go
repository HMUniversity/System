package controllers

import (
	"github.com/HMUniversity/System/controllers/member"
	"github.com/gofiber/fiber/v2"
)

func RegisterControllers(router fiber.Router) {
	c := map[string]Controller{
		"/member": &member.Controller{},
	}
	for k, v := range c {
		v.RegisterRouter(router.Group(k))
	}

}

type Controller interface {
	RegisterRouter(fiber.Router)
}
