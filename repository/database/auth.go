package database

import (
	"demonstrasi/config"
	"demonstrasi/models"
)

func Login(username string) (user models.TbLogin, err error) {
	if err := config.DB.Where("name = ?", username).First(&user).Error; err != nil {
		return models.TbLogin{}, err
	}
	return
}

func Register(user *models.TbLogin) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
