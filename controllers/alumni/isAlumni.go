package alumni

import (
	"github.com/HMUniversity/System/models/alumni_mod"
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/HMUniversity/System/utils"
	"github.com/gofiber/fiber/v2"
)

func isAlumni(ctx *fiber.Ctx) error {
	var req alumni_mod.IsAlumniRequest
	if err := ctx.BodyParser(&req); err != nil {
		err_handler.HandleError(err)
		return utils.BadRequest(ctx, "Cannot load request")
	}

	if req.Username == "" {
		return utils.BadRequest(ctx, "Username is required")
	}

	isAlumniVal, err := config.Get().GitHub.IsInOrg(config.Get().Organisation, req.Username)
	if err != nil {
		return utils.BadRequest(ctx, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(alumni_mod.IsAlumniResponse{
		IsAlumni: isAlumniVal,
	})
}
