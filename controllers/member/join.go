package member

import (
	"github.com/HMUniversity/System/models/member_mod"
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/HMUniversity/System/utils"
	"github.com/gofiber/fiber/v2"
)

func join(ctx *fiber.Ctx) error {
	var req member_mod.JoinRequest
	if err := ctx.BodyParser(&req); err != nil {
		err_handler.HandleError(err)
		return utils.BadRequest(ctx, "Cannot load request")
	}
	user := utils.DeRefString(req.Username)
	email := utils.DeRefString(req.Email)
	if user == "" && email == "" {
		return utils.BadRequest(ctx, "Username or email is required")
	}
	var err error
	if user != "" {
		err = config.Get().GitHub.InviteByUsername(user, config.Get().Organisation)
	} else {
		err = config.Get().GitHub.InviteByEmail(email, config.Get().Organisation)
	}
	if err != nil {
		return utils.BadRequest(ctx, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(member_mod.JoinResponse{
		Message: "Invitation sent",
	})
}
