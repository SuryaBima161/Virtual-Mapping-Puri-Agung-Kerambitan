package controllers

import (
	"demonstrasi/models/payload"
	"demonstrasi/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func CreateGalery(c echo.Context) error {
// 	// Get userID from authenticated user
// 	userID := middlewares.GetUserLoginId(c)
// 	// Parse request body
// 	req := new(payload.AddToCart)
// 	if err := c.Bind(req); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
// 	}

// 	// Validate request
// 	if err := c.Validate(req); err != nil {
// 		return err
// 	}

// 	// Add to cart
// 	err := usecase.AddToCart(userID, req.ProductID, req.Quantity)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "product stock is not available") {
// 			return echo.NewHTTPError(http.StatusBadRequest, "Product stock is not available")
// 		}
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to add to cart")
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Product added to cart successfully",
// 	})
// }

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
