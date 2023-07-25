package dtos

type UserDTO struct {
	Username string `json:"username" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Mobile   string `json:"mobile" validate:"required"`
	Age      int    `json:"age" validate:"required,gt=0"`
	Email    string `json:"email" validate:"required,email"`
}
