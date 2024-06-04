package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateAdminController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := payload.UpdateAdmin{}
	c.Bind(&req)
	if err := usecase.UpdateAdmin(uint(id), req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin",
	})
}
