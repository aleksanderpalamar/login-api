package handlers

import (
	"net/http"

	"github.com/aleksanderpalamar/login-api/db"
	"github.com/aleksanderpalamar/login-api/models"
	"github.com/aleksanderpalamar/login-api/utils"
	"github.com/labstack/echo/v4"
)

var users = map[string]models.User{}

func SignUp(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if result := db.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User created successfully"})
}

func SignIn(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var storedUser models.User
	if result := db.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&storedUser); result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	if storedUser.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	token, err := utils.GenerateToken(storedUser.Username, storedUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func SignOut(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "User signed out successfully"})
}
