package morsencoder

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type MorseEncoderService interface {
	Encode(textToEncode string) (string, error)
}
type Logger interface {
	Debugf(template string, args ...interface{})
}
type Handler struct {
	service MorseEncoderService
	logger  Logger
}

func NewHandler(service MorseEncoderService, logger Logger) *Handler {
	if logger == nil || service == nil {
		return nil
	}
	return &Handler{
		service: service,
		logger:  logger,
	}
}
func (h *Handler) Encode(ctx *fiber.Ctx) error {
	textToEncode := ctx.Query("text", "")
	h.logger.Debugf("Encoding request arrived for the string %s", textToEncode)
	encodedText, err := h.service.Encode(textToEncode)
	if err != nil {
		if errors.Is(err, InvalidTextToEncode) {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).SendString(encodedText)
}
func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/encode", h.Encode)
}
