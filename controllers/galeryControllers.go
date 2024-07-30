package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)
func CreateGalery(c echo.Context) error {
	var req payload.AddGalery
	c.Bind(&req)
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := usecase.CreateGalery(&req, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Galery",
	})
}

func GetGalery(c echo.Context) error {
	var galery []payload.GetGaleryRespone
	galery, err := usecase.GetGalery()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all galery",
		"data":    galery,
	})
}
func GetHomePage(c echo.Context) error {
	galery, err := usecase.GetGaleryByRating()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all galery",
		"data":    galery,
	})
}

func UpdateGalery(c echo.Context) error {
	// Bind the payload
	var inf payload.UpdateGalery
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

	if err := usecase.UpdateGalery(uuid.FromStringOrNil(id), &inf, image); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update galery",
	})
}

func DeleteGalery(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteGalery(uuid.FromStringOrNil(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete galery",
	})
}
