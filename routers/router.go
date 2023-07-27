package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // Import gin-swagger middleware
	"github.com/user/test_template/db"
	"github.com/user/test_template/handlers" // Import the auto-generated docs package
	"github.com/user/test_template/logger"
)

func RunServer() {
	logger.InitLogger("default", 1, 3, 7)
	db.InitialDbConnection()
	router := gin.Default()
	// Register the Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users/getAllUsers", handlers.GetAllUsers)
	router.POST("/users/createUser", handlers.CreateUser)
	router.PUT("/users/updateUserById/:id", handlers.UpdateUser)
	router.GET("/users/getUserById/:id", handlers.GetUserByID)
	router.DELETE("/users/deleteUserById/:id", handlers.DeleteUserByID)
	router.Run("localhost:8080")
}
