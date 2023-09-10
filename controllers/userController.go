package controllers

import (
	"backjeep/models"
	"backjeep/repositories"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.Repo.GetAllUsers()
	if err != nil {
		log.Printf("Error al obtener usuarios: %v", err) // Aquí agregas el log
		return echo.NewHTTPError(http.StatusInternalServerError, "No se pudieron recuperar los usuarios")
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Datos de entrada inválidos")
	}

	user, err := uc.Repo.CreateUser(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "No se pudo crear el usuario")
	}
	return c.JSON(http.StatusCreated, user)
}
