package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/controllers"
)

func LoginRoutes(inRoutes *gin.RouterGroup) {

	inRoutes.POST("/register", controllers.CreateUser)

	inRoutes.POST("/login", controllers.Login)
}
