package routes

import (
	"apiuser/app/models/user"

	uc "apiuser/app/controllers/user_controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := uc.NewUserController(userService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userController.Index)
		v1.GET("/users/:id", userController.GetByID)
		v1.POST("/users", userController.Create)
		v1.PATCH("/users/:id", userController.Update)
		v1.DELETE("/users", userController.Delete)
	}
}
