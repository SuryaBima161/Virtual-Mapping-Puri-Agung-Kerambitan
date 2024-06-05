package payload

type AddComment struct {
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
}
type GetCommentRespone struct {
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
}
