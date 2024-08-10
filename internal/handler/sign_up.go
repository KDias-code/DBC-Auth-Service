package handler

import (
	"auth-service/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	request := c.Body()
	signModel := new(model.SignUp)

	err := json.Unmarshal(request, signModel)
	if err != nil {
		h.logger.Error("failed to unmarshal request body", "error", err)
		return c.Status(400).JSON("invalid request body")
	}

	err = h.signUp.SignUp(*signModel)
	if err != nil {
		h.logger.Error("failed to create user", "error", err)
		return c.Status(500).JSON("failed to create user")
	}

	return c.Status(200).JSON("User success created!")
}
