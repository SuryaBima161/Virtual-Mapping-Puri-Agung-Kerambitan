package database

import (
	"demonstrasi/config"
	"demonstrasi/models"
)

func GetAdminById(id uint) (resp *models.TbLogin, err error) {
	if err := config.DB.Where("user_id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func GetAdminAll(req *models.TbLogin) ([]models.TbLogin, error) {
	var admin []models.TbLogin
	if err := config.DB.Model(req).Find(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
