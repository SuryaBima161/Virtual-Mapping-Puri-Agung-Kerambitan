package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func CreateComment(c echo.Context) error {
	var req payload.AddComment
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := usecase.CreateComment(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create commentar",
	})
}

func DeleteComment(c echo.Context) error {
	id := c.Param("id")
	if err := usecase.DeleteComment(uuid.FromStringOrNil(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete commentar",
	})
}

func GetComment(c echo.Context) error {
	var comment []payload.GetCommentRespone
	comment, err := usecase.GetComment()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all comments",
		"products": comment,
	})
}

func GetCommentById(c echo.Context) error {
	id := c.Param("id")
	comment, err := usecase.GetCommentById(uuid.FromStringOrNil(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get comment by id",
		"data":    comment,
	})

}
