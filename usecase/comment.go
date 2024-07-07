package usecase

import (
	"demonstrasi/models"
	"demonstrasi/models/payload"
	"demonstrasi/repository/database"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func CreateComment(req *payload.AddComment) error {
	comment := models.TbComment{
		Id_Galery: req.Id_Galery,
		Name:      req.Name,
		Comment:   req.Comment,
		Rating:    req.Rating,
	}
	if err := database.CreateComment(&comment); err != nil {
		return err
	}
	return nil
}

func GetComment(status string) ([]payload.GetCommentRespone, error) {
	comments, err := database.GetComment(status)
	if err != nil {
		return nil, err
	}

	var responses []payload.GetCommentRespone
	for _, data := range comments {
		responses = append(responses, payload.GetCommentRespone{
			Name:    data.Name,
			Comment: data.Comment,
			Rating:  data.Rating,
			Reply:   data.ReplyComment,
		})
	}

	return responses, nil
}

func GetCommentById(id uuid.UUID) (resp payload.GetCommentRespone, err error) {
	comment, err := database.GetCommentById(id)
	if err != nil {
		return payload.GetCommentRespone{}, err
	}
	resp = payload.GetCommentRespone{
		Name:    comment.Name,
		Comment: comment.Comment,
	}
	return
}

func DeleteComment(id uuid.UUID) (err error) {
	if _, err := database.GetCommentById(id); err != nil {
		return err
	}
	err = database.DeleteComment(id)
	if err != nil {
		fmt.Println("Delete: error deleting commentar, err:", err)
		return err
	}
	return nil
}

func UpdateReplyComment(id uuid.UUID, req *payload.ReplyCommentRequest) (err error) {
	if _, err := database.GetCommentById(id); err != nil {
		return err
	}
	inf := models.TbComment{
		ReplyComment: req.ReplyComment,
	}
	if err := database.UpdateReplyComment(id, &inf); err != nil {
		return err
	}
	return nil

}

func ValidateComment(id uuid.UUID, req *payload.ValidateComment) error {
	comment, err := database.GetCommentById(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve comment: %v", err)
	}
	comment.Status = "validated"

	if err := database.ValidateComment(id, &comment); err != nil {
		return fmt.Errorf("failed to validate comment: %v", err)
	}

	return nil
}
