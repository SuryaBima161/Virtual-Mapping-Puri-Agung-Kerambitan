package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
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
		"message": "success create monument",
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

func GetMonumentById(c echo.Context) error {
	id := c.Param("id")
	mon, err := usecase.GetMonumentById(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get monument by id",
		"data":    mon,
	})

}

func UpdateMonument(c echo.Context) error {
	var inf payload.UpdateMonument
	id := c.Param("id")
	if err := c.Bind(&inf); err != nil {
		return err
	}
	if err := c.Validate(inf); err != nil {
		return err
	}

	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to read image file: "+err.Error())
	}

	if err := usecase.UpdateMonument(uuid.FromStringOrNil(id), &inf, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update monument",
	})
}
func DeleteMonument(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteMonument(uuid.FromStringOrNil(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete monument",
	})
}
