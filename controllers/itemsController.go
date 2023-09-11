package controllers

import (
	"backjeep/models"
	"backjeep/repositories"
	"backjeep/utils"
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

func CreateItem(c echo.Context) error {
	var req models.CreateItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid data",
		})
	}

	// Subir las imágenes a Cloudinary y obtener las URLs
	imageURLs, err := utils.UploadImages(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to upload image",
		})
	}

	itemRequest := models.CreateItemRequest{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Quantity:    req.Quantity,
		Category:    req.Category,
		Subcategory: req.Subcategory,
		Slug:        req.Slug,
		Store:       req.Store,
		Link:        req.Link,
	}
	itemRequest.Images = imageURLs

	item, err := itemRepo.CreateItem(itemRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create item",
		})
	}

	return c.JSON(http.StatusOK, item)
}
