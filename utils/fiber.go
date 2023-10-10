package utils

import (
	"github.com/HMUniversity/System/models/common_mod"
	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(common_mod.ErrorMessage{
		Message: msg,
	})
}
