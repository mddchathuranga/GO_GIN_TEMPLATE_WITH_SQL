package services

import (
	//"os/user"

	"errors"

	"github.com/user/test_template/dtos"
	"github.com/user/test_template/models"
	"github.com/user/test_template/repositories"
)

func GetAllUsers() []dtos.UserGetDTO {
	var users []models.User
	users = repositories.GetAllUsers()
	usersGetDTO := MapToUserGetDTO(users)
	return usersGetDTO
}

func CreateUser(newDTO dtos.UserDTO) string {
	newUser := MapToUser(newDTO)
	if err := repositories.CreateUser(newUser); err != nil {
		return err.Error()
	}
	return "user Created"
}

func UpdateUser(id int, newDTO dtos.UserDTO) string {

	newUser := MapToUser(newDTO)
	if err := repositories.UpdateUser(id, newUser); err != nil {

		return err.Error()

	}
	return "user updated"

}
func GetUserByID(id int) (dtos.UserGetDTO, error) {
	var user models.User
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return dtos.UserGetDTO{}, errors.New(err.Error())
	}
	userGetDTO := MapToUserGetByIDDTO(user)
	return userGetDTO, nil

}
func DeleteUserByID(id int) string {
	if err := repositories.DeleteUserByID(id); err != nil {
		return err.Error()
	}
	return "deleted"
}

func MapToUser(userDTO dtos.UserDTO) models.User {
	return models.User{
		Username: userDTO.Username,
		Address:  userDTO.Address,
		Mobile:   userDTO.Mobile,
		Age:      userDTO.Age,
		Email:    userDTO.Email,
	}
}

func MapToUserGetDTO(users []models.User) []dtos.UserGetDTO {
	var usersDTO []dtos.UserGetDTO
	for _, user := range users {
		userDTO := dtos.UserGetDTO{
			ID:       user.ID,
			Username: user.Username,
			Address:  user.Address,
			Mobile:   user.Mobile,
			Age:      user.Age,
			Email:    user.Email,
		}
		usersDTO = append(usersDTO, userDTO)
	}
	return usersDTO
}

func MapToUserGetByIDDTO(user models.User) dtos.UserGetDTO {

	userDTO := dtos.UserGetDTO{
		ID:       user.ID,
		Username: user.Username,
		Address:  user.Address,
		Mobile:   user.Mobile,
		Age:      user.Age,
		Email:    user.Email,
	}
	return userDTO
}
