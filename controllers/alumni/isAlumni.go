package alumni

import (
	"github.com/HMUniversity/System/models/alumni_mod"
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/gofiber/fiber/v2"
)

func isAlumni(ctx *fiber.Ctx) error {
	var req alumni_mod.IsAlumniRequest
	if err := ctx.JSON(&req); err != nil {
		err_handler.HandleError(err)
		ctx.Status(fiber.StatusBadRequest).JSON(alumni_mod.IsAlumniResponse{
			Message: "Cannot load request",
		})
		return nil
	}

	if req.Username == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(alumni_mod.IsAlumniResponse{
			Message: "Username is required",
		})
		return nil
	}

	isAlumniVal, err := config.Get().GitHub.IsInOrg(config.Get().Organisation, req.Username)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(alumni_mod.IsAlumniResponse{
			Message: err.Error(),
		})
		return nil
	}
	ctx.Status(fiber.StatusOK).JSON(alumni_mod.IsAlumniResponse{
		IsAlumni: isAlumniVal,
	})
	return nil
}
