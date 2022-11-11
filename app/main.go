package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/vitostamatti/notes-app-golang/app/middlewares"
	"github.com/vitostamatti/notes-app-golang/app/models"
	"github.com/vitostamatti/notes-app-golang/app/routers"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	apiRouter := router.Group("/api")

	routers.LoginRoutes(apiRouter)

	usersRouter := apiRouter.Group("/users")
	usersRouter.Use(middlewares.JwtAuthMiddleware())
	routers.UsersRoutes(usersRouter)

	notesRouter := apiRouter.Group("/notes")
	notesRouter.Use(middlewares.JwtAuthMiddleware())
	routers.NotesRoutes(notesRouter)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000"
	}
	router.Run(":" + port)

}
