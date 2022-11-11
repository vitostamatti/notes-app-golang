package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/controllers"
)

func NotesRoutes(inRoutes *gin.RouterGroup) {

	inRoutes.GET("", controllers.GetNotes)

	inRoutes.POST("", controllers.CreateNote)

	inRoutes.GET("/:id", controllers.GetNoteById)

	inRoutes.PUT("/:id", controllers.UpdateNote)

	inRoutes.DELETE("/:id", controllers.DeleteNote)
}
