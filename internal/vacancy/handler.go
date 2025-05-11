package vacancy

import "github.com/gofiber/fiber/v2"

type VacancyHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &VacancyHandler{
		router: router,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	return c.SendString("Vacancy")
}
