package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/controllers"
)

func UsersRoutes(router *gin.RouterGroup) {

	router.GET("", controllers.GetUsers)

	router.POST("", controllers.CreateUser)

	router.GET("/me", controllers.GetCurrentUser)

	router.GET("/:id", controllers.GetUserById)

	router.PUT("/:id", controllers.UpdateUser)

	router.DELETE("/:id", controllers.DeleteUser)
}
