package alumni

import (
	"github.com/HMUniversity/System/models/alumni_mod"
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/gofiber/fiber/v2"
)

func isAlumni(ctx *fiber.Ctx) error {
	var req alumni_mod.IsAlumniRequest
	if err := ctx.BodyParser(&req); err != nil {
		err_handler.HandleError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(alumni_mod.IsAlumniResponse{
			Message: "Cannot load request",
		})
	}

	if req.Username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(alumni_mod.IsAlumniResponse{
			Message: "Username is required",
		})
	}

	isAlumniVal, err := config.Get().GitHub.IsInOrg(config.Get().Organisation, req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(alumni_mod.IsAlumniResponse{
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(alumni_mod.IsAlumniResponse{
		IsAlumni: isAlumniVal,
	})
}
