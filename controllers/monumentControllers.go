package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMonument(c echo.Context) error {
	var req payload.MonumentRequest
	c.Bind(&req)
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := usecase.CreateMonument(&req, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Monument",
	})
}

func GetMonument(c echo.Context) error {
	var monument []payload.GetMonumentRespone
	monument, err := usecase.GetMonument()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all monument",
		"data":    monument,
	})
}
