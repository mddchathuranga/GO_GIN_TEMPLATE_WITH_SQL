package dtos

type PostGetDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
}
