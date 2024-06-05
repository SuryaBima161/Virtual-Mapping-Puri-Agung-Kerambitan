package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func CreateInformation(c echo.Context) error {
	var req payload.AddInformation
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := usecase.CreateInformation(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create information",
	})
}

func UpdateInfortmation(c echo.Context) error {
	var inf payload.UpdateInfortmation
	id := c.Param("id")
	c.Bind(&inf)
	if err := c.Validate(inf); err != nil {
		return err
	}
	if err := usecase.UpdateInfortmation(uuid.FromStringOrNil(id), &inf); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update information",
	})
}

func DeleteInformation(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteInfortmation(uuid.FromStringOrNil(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete information",
	})
}

func GetInformation(c echo.Context) error {
	var inf []payload.GetInformationRespone
	inf, err := usecase.GetInformation()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all information",
		"products": inf,
	})
}

func GetInformationById(c echo.Context) error {
	id := c.Param("id")
	inf, err := usecase.GetInformationById(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get information by id",
		"data":    inf,
	})

}
