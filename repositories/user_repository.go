package repositories

import (
	"errors"

	"github.com/user/test_template/db"
	"github.com/user/test_template/logger"
	"github.com/user/test_template/models"
)

func GetAllUsers() []models.User {
	var users []models.User
	db.DB.Find(&users)
	return users
}

func CreateUser(user models.User) string {
	logger := logger.GetLogger() // Get the initialized logger instance
	result := db.DB.Create(&user)
	if result.Error != nil {
		logger.Info("user cannot create :", result.Error)
		return "operation failed"

	}
	return "user saved"
}

func UpdateUser(id int, updateUser models.User) error {
	logger := logger.GetLogger() // Get the initialized logger instance
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		logger.Info("user not found")
		return errors.New("user not found")
	}
	result = db.DB.Model(&user).Updates(updateUser)

	if result.Error != nil {
		logger.Info("user cannot create :", result.Error)
		return errors.New(result.Error.Error())
	}
	return nil
}
func GetUserByID(id int) (models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}
	return user, nil

}

func DeleteUserByID(id int) error {

	logger := logger.GetLogger() // Get the initialized logger instance
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		logger.Info("user not found")
		return errors.New("user not found")
	}
	db.DB.Delete(&user, id)
	return nil
}
