package controllers

import (
	"backjeep/models"
	"backjeep/repositories"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	e := echo.New()
	mockUserController := &UserController{
		Repo: &MockUserRepository{},
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	mockUser := models.User{
		ID:       1,
		Name:     "Shaggy",
		Email:    "shaggy@doo.com",
		Password: string(hashedPassword),
	}

	t.Run("successful login", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "shaggy@doo.com",
			"password": "secret",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		repositories.MockUserDB["shaggy@doo.com"] = &mockUser

		if assert.NoError(t, mockUserController.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("bind failure", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.Error(t, mockUserController.Login(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "notfound@dubby.com",
			"password": "secret",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.Error(t, mockUserController.Login(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("wrong password", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "shaggy@doo.com",
			"password": "wrongpassword",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		repositories.MockUserDB["shaggy@doo.com"] = &mockUser

		if assert.Error(t, mockUserController.Login(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})
}
