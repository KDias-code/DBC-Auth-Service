package handler

import (
	"auth-service/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignIn(c *fiber.Ctx) error {
	request := c.Body()
	signModel := new(model.SignIn)

	err := json.Unmarshal(request, signModel)
	if err != nil {
		h.logger.Error("failed to unmarshal request into the model", "error", err)
		return c.Status(400).JSON("invalid request body")
	}

	token, err := h.signIn.SignIn(*signModel)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).SendString(token)
}
