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
	if err := ctx.JSON(&req); err != nil {
		err_handler.HandleError(err)
		ctx.Status(fiber.StatusBadRequest).JSON(member_mod.JoinResponse{
			Message: "Cannot load request",
		})
		return nil
	}
	user := utils.DeRefString(req.Username)
	email := utils.DeRefString(req.Email)
	if user == "" && email == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(member_mod.JoinResponse{
			Message: "Username or email is required",
		})
		return nil
	}
	var err error
	if user != "" {
		err = config.Get().GitHub.InviteByUsername(user, config.Get().Organisation)
	} else {
		err = config.Get().GitHub.InviteByEmail(email, config.Get().Organisation)
	}
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(member_mod.JoinResponse{
			Message: err.Error(),
		})
		return nil
	}
	ctx.Status(fiber.StatusOK).JSON(member_mod.JoinResponse{
		Message: "Invitation sent",
	})
	return nil
}
