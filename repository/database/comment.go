package database

import (
	"demonstrasi/config"
	"demonstrasi/models"

	uuid "github.com/satori/go.uuid"
)

func CreateComment(req *models.TbComment) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func GetCommentValidated(status string) ([]models.TbComment, error) {
	var comments []models.TbComment
	db := config.DB
	if status != "" {
		db = db.Where("status = ?", status)
	}

	if err := db.Find(&comments).Error; err != nil {
		return comments, err
	}

	return comments, nil
}
func GetComment() ([]models.TbComment, error) {
	var comments []models.TbComment
	db := config.DB
	if err := db.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentById(id uuid.UUID) (resp models.TbComment, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func DeleteComment(id uuid.UUID) error {
	inf := models.TbComment{}
	if err := config.DB.Model(&inf).Where("id = ?", id).Delete(&inf).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReplyComment(id uuid.UUID, inf *models.TbComment) error {
	if err := config.DB.Model(inf).Where("id = ?", id).Updates(inf).Error; err != nil {
		return err
	}
	return nil
}

func ValidateComment(id uuid.UUID, inf *models.TbComment) error {
	if err := config.DB.Model(inf).Where("id = ?", id).Updates(inf).Error; err != nil {
		return err
	}
	return nil
}
