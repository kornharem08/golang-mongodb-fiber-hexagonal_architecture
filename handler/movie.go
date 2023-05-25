package handler

import (
	"basic/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type movieHandler struct {
	movSrv service.MovieService
}

func NewMovieHandler(movSrv service.MovieService) movieHandler {
	return movieHandler{movSrv: movSrv}
}

func (h movieHandler) GetMovies(c *fiber.Ctx) error {
	pageStr := c.Query("page") // Get the value of the "page" query parameter
	limitStr := c.Query("limit")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid value for 'page'",
		})
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid value for 'limit'",
		})
	}

	responses, err := h.movSrv.GetMovies(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unexpected error",
		})
	}

	return c.JSON(responses)
}
