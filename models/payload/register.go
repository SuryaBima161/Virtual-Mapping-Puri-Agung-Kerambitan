package payload

type Register struct {
	Username       string `json:"username" form:"username" validate:"required"`
	Password       string `json:"password" form:"password" validate:"required,min=8"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required"`
}
