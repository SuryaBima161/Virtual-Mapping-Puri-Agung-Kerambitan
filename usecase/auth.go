package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(req *payload.GetAdmin) (models.TbLogin, error) {
	user, err := database.Login(req.Username)
	if err != nil {
		return models.TbLogin{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.TbLogin{}, err
	}
	return user, nil
}

func Register(req *payload.Register) (models.TbLogin, error) {
	if req.Password != req.RetypePassword {
		return models.TbLogin{}, echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.TbLogin{}, err
	}
	user := models.TbLogin{
		Username: req.Username,
		Password: string(password),
	}
	if err := database.Register(&user); err != nil {
		return models.TbLogin{}, err
	}
	return user, nil
}
