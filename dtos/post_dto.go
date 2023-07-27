package dtos

type PostDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  int    `json:"user_id"`
}
