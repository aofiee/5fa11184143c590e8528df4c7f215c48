package handler

import (
	"count-beef/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type (
	Handler struct {
		Service ports.Service
	}
)

func NewHandler(service ports.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (hdl *Handler) GetSummary(c *fiber.Ctx) error {
	o, err := hdl.Service.GetSummary()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	log.Println("Response:", o)
	return c.JSON(o)
}
