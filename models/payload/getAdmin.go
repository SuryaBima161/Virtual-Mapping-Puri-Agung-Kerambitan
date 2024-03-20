package payload

type GetAdmin struct {
	// ID       uuid.UUID `json:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateAdmin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
