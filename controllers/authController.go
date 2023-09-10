package controllers

import (
	"backjeep/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserController) Login(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Datos de entrada inválidos")
	}

	storedUser, err := uc.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = storedUser.ID
	claims["email"] = storedUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	jwtSecret := viper.GetString("JWT_SECRET")
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login exitoso!",
	})
}
