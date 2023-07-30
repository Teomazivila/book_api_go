package main

import (
	"bookapi/config"
	"bookapi/controller"
	"bookapi/entity"
	"bookapi/middleware"

	"github.com/gin-gonic/gin"
)

var Users []entity.User

func main() {
	config.ConnectDB()

	defer config.CloseDb()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
		}

		book := v1.Group("/book")
		{
			book.GET("/", controller.GetAllBooks)
			book.GET("/:id", controller.GetBook)
			book.POST("/", middleware.Authorized(), controller.InsertBook)
			book.PUT("/:id", controller.UpdateBook) // DTO -> AUTH + OWNER
			book.DELETE("/:id", middleware.Authorized(), controller.DeleteBook)
		}

		user := v1.Group("/user")
		{
			user.GET("/", controller.GetAllUsers) // SEM AUTH -> DTO RESPONSE (ID, NAME, EMAIL)
			user.POST("/", controller.Register)
			user.GET("/:id", controller.Profile)          // AUTH - OWNER -> DTO RESPONSE (ID, NAME, EMAIL, PROFILE PICTURE)
			user.PUT("/:id", controller.UpdateProfile)    // AUTH - OWNER
			user.DELETE("/:id", controller.DeleteAccount) // AUTH - OWNER
		}

	}
	router.Run(":3000")
}
