package payload

import uuid "github.com/satori/go.uuid"

type MonumentRequest struct {
	InformationID uuid.UUID `json:"information_id" form:"information_id"`
	Image         string    `json:"image" form:"image"`
}

type GetMonumentRespone struct {
	Image string `json:"image" form:"image"`
}
