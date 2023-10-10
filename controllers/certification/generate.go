package certification

import (
	"github.com/HMUniversity/System/models/certification_mod"
	"github.com/HMUniversity/System/models/member_mod"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/HMUniversity/System/modules/pdf_gen"
	"github.com/gofiber/fiber/v2"
	"time"
)

func getCurrentDate() string {
	return time.Now().Format("31/12/2006")
}

func getRandomPath() string {
	return ""
}

func generate(ctx *fiber.Ctx) error {
	var req certification_mod.GenerationRequest
	if err := ctx.BodyParser(&req); err != nil {
		err_handler.HandleError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(member_mod.JoinResponse{
			Message: "Cannot load request",
		})
	}
	data := pdf_gen.CertificationData{
		Programme: req.Programme,
		Date:      getCurrentDate(),
		Receiver:  req.Receiver,
		Degree:    req.Degree,
	}
	path := getRandomPath()
	err := pdf_gen.NewCertification(data, path)
	if err != nil {
		err_handler.HandleError(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(member_mod.JoinResponse{
			Message: "gen failed",
		})
	}
	return ctx.Status(fiber.StatusOK).SendFile(path)
}
