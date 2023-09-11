package controllers

import (
	"backjeep/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockUsers = []models.User{
		{ID: 1, Name: "Pablo Marmol", Email: "pablo@marmol.com"},
	}
	createUserJSON = `{"name":"Pablo Marmol","email":"pablo@marmol.com","password":"SonLosPicapiedras"}`
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetAllUsers() ([]models.User, error) {
	return mockUsers, nil
}

func (m *MockUserRepository) CreateUser(user models.CreateUserRequest) (models.User, error) {
	return mockUsers[0], nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (models.User, error) {
	for _, user := range mockUsers {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, fmt.Errorf("No user found with email %s", email)
}

func TestGetAllUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := &UserController{
		Repo: &MockUserRepository{},
	}

	if assert.NoError(t, userController.GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Pablo Marmol")
	}
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userController := &UserController{
		Repo: &MockUserRepository{},
	}

	if assert.NoError(t, userController.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Pablo Marmol")
	}
}
