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
		Name:    req.Name,
		Comment: req.Comment,
	}
	if err := database.CreateComment(&comment); err != nil {
		return err
	}
	return nil
}

func GetComment() (resp []payload.GetCommentRespone, err error) {
	comment, err := database.GetComment()
	if err != nil {
		return []payload.GetCommentRespone{}, err
	}
	resp = make([]payload.GetCommentRespone, len(comment))
	for i, data := range comment {
		resp[i] = payload.GetCommentRespone{
			Name:    data.Name,
			Comment: data.Comment,
		}
	}

	return resp, nil
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
