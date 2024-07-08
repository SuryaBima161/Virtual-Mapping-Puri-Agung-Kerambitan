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

func GetCommentValidated(c echo.Context) error {
	status := "validated"
	comments, err := usecase.GetCommentValidated(status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all validated comments",
		"data":    comments,
	})
}

func GetComment(c echo.Context) error {
	comments, err := usecase.GetComment()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all comments",
		"data":    comments,
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

func UpdateReplyComment(c echo.Context) error {
	var inf payload.ReplyCommentRequest
	id := c.Param("id")
	c.Bind(&inf)
	if err := c.Validate(inf); err != nil {
		return err
	}
	if err := usecase.UpdateReplyComment(uuid.FromStringOrNil(id), &inf); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success reply comment",
	})
}

func ValideteComment(c echo.Context) error {
	var inf payload.ValidateComment
	id := c.Param("id")
	c.Bind(&inf)
	if err := c.Validate(inf); err != nil {
		return err
	}
	if err := usecase.ValidateComment(uuid.FromStringOrNil(id), &inf); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update status",
	})
}
