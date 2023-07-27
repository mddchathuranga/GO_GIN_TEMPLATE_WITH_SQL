package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user/test_template/dtos"
	"github.com/user/test_template/exutilities"
	"github.com/user/test_template/logger"
	"github.com/user/test_template/services"
)

// @Tags Post
// CreatePost creates a new post.
// @Summary Create a new post
// @Description Create a new Post with the given details
// @ID create-post
// @Accept json0
// @Produce json
// @Param post body dtos.PostDTO true "Post object that needs to be created"
// @Success 201 {object} dtos.PostDTO
// @Failure 400 {object} exutilities.ErrorResponse
// @Router /posts/createpost/{id} [post]
func CreatePost(c *gin.Context) {
	logger := logger.GetLogger()
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var newPostDTO dtos.PostDTO
	if err := c.BindJSON(&newPostDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info(err)
		return
	}
	newPostDTO.UserID = id
	logger.Info("Received post create request:", newPostDTO)
	if err := exutilities.ValidatePostDTO(newPostDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info(" validation failed:", err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, services.CreatePost(newPostDTO))

	logger.Info("Post  created for user id:" + idStr)
}

// @Tags Post
// UpdatePost update by ID.
// @Summary Update Post By ID
// @Description Update an existing Post with the given ID
// @ID Update-post
// @Accept json
// @Produce json
// @Param id path int true "Post ID to be updated"
// @Param post body dtos.PostDTO true "Update Post Object"
// @Success 200 {object} dtos.PostDTO
// @Failure 400 {object} exutilities.ErrorResponse
// @Router /posts/updatePost/{id} [put]
func UpdatePost(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var updatePostDTO dtos.PostDTO
	if err := c.BindJSON(&updatePostDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info(err)
		return
	}
	logger.Info("Received update post reqest with id:", updatePostDTO)
	if err := exutilities.ValidatePostDTO(updatePostDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info(" validation failed:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, services.UpdatePost(id, updatePostDTO))

}

// @Tags Post
// GetPostByID gets a Post by ID.
// @Summary Get a Post by ID
// @Description Get a Post with the given ID
// @ID get-post-by-id
// @Accept json
// @Produce json
// @Param id path int true "User ID to retrieve"
// @Success 200 {object} dtos.PostDTO
// @Failure 400 {object} exutilities.ErrorResponse
// @Router /posts/getPostById/{id} [get]
func GetPostByID(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var postGetDTO, err = services.GetPostByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info(err.Error())
		return
	}
	logger.Info("Received get user details by ID :"+idStr, postGetDTO)
	c.IndentedJSON(http.StatusOK, postGetDTO)
}

// @Tags Post
// @Summary Delete a post by ID
// @Description Delete a post with the given ID
// @ID delete-post-by-id
// @Produce json
// @Param id path int true "Post ID to delete"
// @Success 204
// @Failure 404 {object} exutilities.ErrorResponse
// @Router /posts/deletePost/{id} [delete]
func DeletePostByID(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	response := services.DeletePostByID(id)
	if response != "deleted" {
		c.IndentedJSON(http.StatusNotFound, exutilities.ErrorResponse{Message: response})
		logger.Info(response)
		return
	}
	logger.Info(response)
	c.IndentedJSON(http.StatusNoContent, response)
}

// @Tags Post
// GetPostByUserID gets a Post by Users ID.
// @Summary Get a Post by Users ID
// @Description Get a Post with the given user ID
// @ID get-post-by--user-id
// @Accept json
// @Produce json
// @Param id path int true "User ID to retrieve"
// @Success 200 {object} dtos.PostGetDTO
// @Failure 400 {object} exutilities.ErrorResponse
// @Router /posts/getPostByUserId/{id} [get]
func GetPostByUserID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var postGetDTO = services.GetPostByUserID(id)
	logger := logger.GetLogger() // Get the initialized logger instance
	logger.Info("Sending post related to user ID  :"+idStr, postGetDTO)
	c.IndentedJSON(http.StatusOK, postGetDTO)
}
