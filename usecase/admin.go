package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
)

func GetAdminAll(req *models.TbLogin) ([]models.TbLogin, error) {
	admins, err := database.GetAdminAll(req)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func GetAdminById(id uint) (resp payload.GetAdmin, err error) {
	admin, err := database.GetAdminById(id)
	if err != nil {
		return payload.GetAdmin{}, err
	}
	resp = payload.GetAdmin{
		Username: admin.Username,
		Password: admin.Password,
	}
	return
}
