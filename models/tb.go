package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TbLogin struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Base struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

type TbInformation struct {
	IDLogin    uint   `json:"id_login" form:"id_login"`
	JudulFoto  string `json:"judul_foto" form:"judul_foto"`
	NamaLokasi string `json:"nama_lokasi" form:"nama_lokasi"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
}

type TbComment struct {
	IDComment uuid.UUID `json:"id_comment" form:"id_comment"`
	Name      string    `json:"name" form:"name"`
	Comment   string    `json:"comment" form:"comment"`
}

type TbGalery struct {
	IDGalery       uuid.UUID     `json:"id_galery" form:"id_galery"`
	TbInformation  TbInformation `json:"tb_information" form:"tb_information"`
	Id_Information uuid.UUID     `json:"id_information" form:"id_information"`
	Image          string        `json:"image" form:"image"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
