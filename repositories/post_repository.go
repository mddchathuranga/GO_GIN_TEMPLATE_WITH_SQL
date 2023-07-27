package repositories

import (
	"errors"

	"github.com/user/test_template/db"
	"github.com/user/test_template/logger"
	"github.com/user/test_template/models"
)

func CreatePost(post models.Post) error {
	logger := logger.GetLogger() // Get the initialized logger instance
	result := db.DB.Create(&post)
	if result.Error != nil {
		logger.Info(result.Error)
		return errors.New(result.Error.Error())

	}
	return nil
}

func GetPostById(id int) (models.Post, error) {
	var post models.Post
	result := db.DB.First(&post, id)
	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}
	return post, nil

}
func UpdatePost(id int, updatePost models.Post) error {
	logger := logger.GetLogger() // Get the initialized logger instance
	var post models.Post
	result := db.DB.First(&post, id)
	if result.Error != nil {
		logger.Info(result.Error)
		return errors.New(result.Error.Error())
	}
	result = db.DB.Model(&post).Updates(updatePost)

	if result.Error != nil {
		logger.Info(result.Error)
		return errors.New(result.Error.Error())
	}
	return nil
}

func DeletePostByID(id int) error {

	logger := logger.GetLogger() // Get the initialized logger instance
	var post models.Post
	result := db.DB.First(&post, id)
	if result.Error != nil {
		logger.Info(result.Error)
		return errors.New(result.Error.Error())
	}
	db.DB.Delete(&post, id)
	return nil
}
func GetPostByUserID(userID int) []models.Post {
	var posts []models.Post
	logger := logger.GetLogger()
	result := db.DB.Where("user_id=?", userID).Find(&posts)
	if result.Error != nil {
		logger.Info(result.Error)
		panic(result.Error)

	}
	return posts

}
