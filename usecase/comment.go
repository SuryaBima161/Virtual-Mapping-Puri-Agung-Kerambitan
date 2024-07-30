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

func GetCommentValidated(status string) ([]payload.GetCommentValidateRespone, error) {
	comments, err := database.GetCommentValidated(status)
	if err != nil {
		return nil, err
	}

	var responses []payload.GetCommentValidateRespone
	for _, data := range comments {
		responses = append(responses, payload.GetCommentValidateRespone{
			Name:    data.Name,
			Comment: data.Comment,
			Rating:  data.Rating,
			Reply:   data.ReplyComment,
		})
	}

	return responses, nil
}

func GetComment() ([]payload.GetCommentRespone, error) {
	comments, err := database.GetComment()
	if err != nil {
		return nil, err
	}

	var responses []payload.GetCommentRespone
	for _, data := range comments {
		galleryResponse := payload.GetGaleryForCommentRespone{
			Id_Galery:   data.TbGalery.ID,
			Judul_foto:  data.TbGalery.TbInformation.JudulFoto,
			Nama_lokasi: data.TbGalery.TbInformation.NamaLokasi,
		}

		commentResponse := payload.GetCommentRespone{
			ID:                         data.ID,
			Name:                       data.Name,
			Comment:                    data.Comment,
			Rating:                     data.Rating,
			Reply:                      data.ReplyComment,
			Status:                     data.Status,
			GetGaleryForCommentRespone: galleryResponse,
		}

		responses = append(responses, commentResponse)
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

func GetCommentByGalleryID(galleryID uuid.UUID) ([]payload.GetCommentRespone, error) {
	comments, err := database.GetCommentByIdGalery(galleryID)
	if err != nil {
		return nil, err
	}

	var responses []payload.GetCommentRespone
	for _, data := range comments {
		galleryResponse := payload.GetGaleryForCommentRespone{
			Id_Galery:   data.TbGalery.ID,
			Judul_foto:  data.TbGalery.TbInformation.JudulFoto,
			Nama_lokasi: data.TbGalery.TbInformation.NamaLokasi,
		}

		commentResponse := payload.GetCommentRespone{
			ID:                         data.ID,
			Name:                       data.Name,
			Comment:                    data.Comment,
			Rating:                     data.Rating,
			Reply:                      data.ReplyComment,
			Status:                     data.Status,
			GetGaleryForCommentRespone: galleryResponse,
		}

		responses = append(responses, commentResponse)
	}

	return responses, nil
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

	if _, err := database.GetCommentById(id); err != nil {
		return err
	}
	inf := models.TbComment{
		Status: "validated",
	}
	if err := database.ValidateComment(id, &inf); err != nil {
		return err
	}
	return nil
}
