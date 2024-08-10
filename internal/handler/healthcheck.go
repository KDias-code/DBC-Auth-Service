package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON("OK")
}
