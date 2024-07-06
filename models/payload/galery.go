package payload

import uuid "github.com/satori/go.uuid"

type AddGalery struct {
	InformationID uuid.UUID `json:"information_id" form:"information_id"`
	Image         string    `json:"image" form:"image"`
}

type UpdateGalery struct {
	Image string `json:"image" form:"image"`
}
type GetGaleryRespone struct {
	Image  string  `json:"image" form:"image"`
	Rating float64 `json:"rating" form:"rating"`
}

type GaleryWithRating struct {
	ID     uuid.UUID `json:"id"`
	Image  string    `json:"image"`
	Rating float64   `json:"rating"`
}
type GetHomePageRespone struct {
	ID     uuid.UUID `json:"id" form:"id"`
	Image  string    `json:"image" form:"image"`
	Rating float64   `json:"rating" form:"rating"`
}
