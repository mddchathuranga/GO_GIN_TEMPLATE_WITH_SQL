package exutilities

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/test_template/dtos"
)

func ValidateUserDTO(userDTO dtos.UserDTO) error {
	validate := validator.New()
	return validate.Struct(userDTO)

}
