package controllers

import (
	"demonstrasi/middlewares"
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var req payload.GetAdmin
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := usecase.Login(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// parseUUID, err := uuid.Parse(user.ID)
	// parseUUID := string(user.ID)

	token, err := middlewares.CreateToken(user.Username, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}

func RegisterController(c echo.Context) error {
	var req payload.Register
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := usecase.Register(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := middlewares.CreateToken(user.Username, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
		"token":   token,
	})
}
