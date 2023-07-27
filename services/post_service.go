package services

import (
	"errors"

	"github.com/user/test_template/db"
	"github.com/user/test_template/dtos"
	"github.com/user/test_template/logger"
	"github.com/user/test_template/models"
	"github.com/user/test_template/repositories"
	"gorm.io/gorm"
)

func CreatePost(newPostDTO dtos.PostDTO) string {
	if !UserExists(newPostDTO.UserID) {
		return "no user found"
	}
	newPost := MapToPost(newPostDTO)
	if err := repositories.CreatePost(newPost); err != nil {
		return err.Error()

	}
	return "post created"

}

func UpdatePost(id int, newDTO dtos.PostDTO) string {

	newPost := MapToPost(newDTO)
	if err := repositories.UpdatePost(id, newPost); err != nil {

		return err.Error()

	}
	return "Post updated"

}

func DeletePostByID(id int) string {
	if err := repositories.DeletePostByID(id); err != nil {
		return err.Error()
	}
	return "deleted"
}

func GetPostByID(id int) (dtos.PostDTO, error) {
	var post models.Post
	post, err := repositories.GetPostById(id)
	if err != nil {
		return dtos.PostDTO{}, errors.New(err.Error())
	}
	postGetDTO := MapToPostGetByIDDTO(post)
	postGetDTO.UserID = int(post.ID)
	return postGetDTO, nil

}

func GetPostByUserID(id int) []dtos.PostGetDTO {
	//var posts []models.Post
	posts := repositories.GetPostByUserID(id)
	postGetDTO := MapToPostGetDTO(posts)
	return postGetDTO
}

func MapToPost(postDTO dtos.PostDTO) models.Post {
	return models.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		UserID:  uint(postDTO.UserID),
	}
}

func UserExists(userID int) bool {
	var user models.User
	logger := logger.GetLogger()
	result := db.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false // User does not exist
		}
		logger.Info(result.Error)
		return false // Error occurred while querying the database
	}
	return true // User exists
}

func MapToPostGetByIDDTO(post models.Post) dtos.PostDTO {

	postDTO := dtos.PostDTO{
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.User.ID,
	}
	return postDTO
}

func MapToPostGetDTO(posts []models.Post) []dtos.PostGetDTO {
	var postsDTO []dtos.PostGetDTO
	for _, post := range posts {
		postDTO := dtos.PostGetDTO{
			Title:   post.Title,
			Content: post.Content,
			PostID:  post.ID,
		}
		postsDTO = append(postsDTO, postDTO)
	}
	return postsDTO
}
