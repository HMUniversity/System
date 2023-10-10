package certification

import (
	"github.com/HMUniversity/System/models/certification_mod"
	"github.com/HMUniversity/System/modules/config"
	"github.com/HMUniversity/System/modules/err_handler"
	"github.com/HMUniversity/System/modules/pdf_gen"
	"github.com/HMUniversity/System/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"path"
	"strings"
	"time"
)

func getCurrentDate() string {
	return time.Now().Format("31/12/2006")
}

func getRandomPath() string {
	return path.Join(config.Get().CertOutput,
		strings.Replace(uuid.New().String(), "-", "", -1)+".pdf")
}

func generate(ctx *fiber.Ctx) error {
	var req certification_mod.GenerationRequest
	if err := ctx.BodyParser(&req); err != nil {
		err_handler.HandleError(err)
		return utils.BadRequest(ctx, "Cannot load request")
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
		return utils.BadRequest(ctx, "gen failed")
	}
	return ctx.Status(fiber.StatusOK).SendFile(path)
}
