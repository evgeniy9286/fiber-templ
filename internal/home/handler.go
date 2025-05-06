package home

import (
	"fiber-templ-learn/pkg/tadapter"
	"fiber-templ-learn/views"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello Api")
	})
	h.router.Get("/", h.home)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main("Vasia")
	return tadapter.Render(c, component)
} 

