package middleware

import (
	"github.com/HMUniversity/System/modules/config"
	"github.com/gofiber/fiber/v2"
)

var ErrNotAuthorized = fiber.NewError(fiber.StatusUnauthorized, "Not Authorized")

func MidAuth(ctx *fiber.Ctx) error {
	if !config.Get().NeedAuth {
		return nil
	}
	if val, ok := ctx.GetReqHeaders()["Authorization"]; ok {
		for _, v := range config.Get().APIKeys {
			if val == v {
				return nil
			}
		}
		return ErrNotAuthorized
	}
	return ErrNotAuthorized
}
