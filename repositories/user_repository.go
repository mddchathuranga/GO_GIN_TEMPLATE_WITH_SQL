package repositories

import (
	"errors"

	"github.com/user/test_template/db"
	"github.com/user/test_template/models"
)

func GetAllUsers() []models.User {

	return db.Users
}

func CreateUser(user models.User) string {
	db.Users = append(db.Users, user)
	return "user saved"
}

func GetLengthOfUsers() int {
	return len(db.Users) + 1
}

func UpdateUser(id int, updateUser models.User) error {
	for i, u := range db.Users {
		if u.ID == id {
			updateUser.ID = id
			db.Users[i] = updateUser
			return nil
		}
	}
	return errors.New("user not found")
}
func GetUserByID(id int) (models.User, error) {
	for i, u := range db.Users {
		if u.ID == id {
			return db.Users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func DeleteUserByID(id int) error {
	for i, u := range db.Users {
		if u.ID == id {
			db.Users = append(db.Users[:i], db.Users[i+1:]...)
			return nil

		}
	}
	return errors.New("user not found")
}
