package controllers

import (
	"backjeep/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

var itemRepo repositories.ItemRepository = &repositories.ItemRepo{}

func GetItems(c echo.Context) error {
	items, err := itemRepo.GetAllItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al obtener los items.",
		})
	}

	return c.JSON(http.StatusOK, items)
}
