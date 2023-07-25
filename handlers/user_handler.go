package handlers

import (
	// Import the Swagger annotations package
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/user/test_template/docs"
	"github.com/user/test_template/dtos"
	"github.com/user/test_template/exutilities"
	"github.com/user/test_template/logger"
	"github.com/user/test_template/services"
)

// @Tags User
// GetAllUsers gets all users.
// @Summary Get all users
// @Description Get a list of all users
// @ID get-all-users
// @Produce json
// @Success 200 {array} dtos.UserGetDTO
// @Router /users/getAllUsers [get]
func GetAllUsers(c *gin.Context) {
	var usersDTO = services.GetAllUsers()
	logger := logger.GetLogger() // Get the initialized logger instance
	logger.Info("Sending all users :", usersDTO)
	c.IndentedJSON(http.StatusOK, usersDTO)
}

// @Tags User
// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user with the given details
// @ID create-user
// @Accept json
// @Produce json
// @Param user body dtos.UserDTO true "User object that needs to be created"
// @Success 201 {object} dtos.UserDTO
// @Failure 400 {object} exutilities.ErrorResponse
// @Router /users/createUser [post]
func CreateUser(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	var newUserDTO dtos.UserDTO
	if err := c.BindJSON(&newUserDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: "user not created"})
		logger.Info("User creation failed:")
		return
	}
	logger.Info("Received new user DTO:", newUserDTO)
	if err := exutilities.ValidateUserDTO(newUserDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, exutilities.ErrorResponse{Message: err.Error()})
		logger.Info("User validation failed:", err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, services.CreateUser(newUserDTO))
	logger.Info("User created:")
}

// @Tags User
// UpdateUser updates a user by ID.
// @Summary Update a user by ID
// @Description Update an existing user with the given ID
// @ID update-user-by-id
// @Accept json
// @Produce json
// @Param id path int true "User ID to be updated"
// @Param user body dtos.UserDTO true "Updated user object"
// @Success 200 {object} dtos.UserDTO
// @Failure 404 {object} exutilities.ErrorResponse
// @Router /users/updateUserById/{id} [put]
func UpdateUser(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var updateUserDTO dtos.UserDTO
	if err := c.BindJSON(&updateUserDTO); err != nil {
		c.IndentedJSON(http.StatusNotFound, exutilities.ErrorResponse{Message: "user not found"})
		logger.Info("user not found under user ID :" + idStr)
		return
	}
	logger.Info("Received update user reqest with :", updateUserDTO)
	c.IndentedJSON(http.StatusOK, services.UpdateUser(id, updateUserDTO))

}

// @Tags User
// GetUserByID gets a user by ID.
// @Summary Get a user by ID
// @Description Get a user with the given ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID to retrieve"
// @Success 200 {object} dtos.UserGetDTO
// @Failure 404 {object} exutilities.ErrorResponse
// @Router /users/getUserById/{id} [get]
func GetUserByID(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var userGetDTO, err = services.GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, exutilities.ErrorResponse{Message: "user not found"})
		logger.Info("user not found under user ID :" + idStr)
		return
	}
	logger.Info("Received get user details by ID :"+idStr, userGetDTO)
	c.IndentedJSON(http.StatusOK, userGetDTO)
}

// @Tags User
// DeleteUserByID deletes a user by ID.
// @Summary Delete a user by ID
// @Description Delete a user with the given ID
// @ID delete-user-by-id
// @Produce json
// @Param id path int true "User ID to delete"
// @Success 204
// @Failure 404 {object} exutilities.ErrorResponse
// @Router /users/deleteUserById/{id} [delete]
func DeleteUserByID(c *gin.Context) {
	logger := logger.GetLogger() // Get the initialized logger instance
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := services.DeleteUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, exutilities.ErrorResponse{Message: "user not found"})
		logger.Info("user not found under user ID :" + idStr)
		return
	}
	logger.Info("Received Delete user details by ID :" + idStr)
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "user deleted"})
}
