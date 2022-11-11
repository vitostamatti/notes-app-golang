package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/models"
)

type NoteRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func GetNotes(c *gin.Context) {
	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}
	notes := models.GetNotesByAuthor(fmt.Sprint(user.ID))
	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	var reqNote NoteRequest
	var resNote models.Note

	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	if err := c.BindJSON(&reqNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resNote.AuthorID = user.ID
	resNote.Name = reqNote.Name
	resNote.Content = reqNote.Content

	createdNote, err := resNote.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdNote)

}

func GetNoteById(c *gin.Context) {
	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	id := c.Param("id")
	note := models.GetNoteById(id)

	if note.AuthorID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	id := c.Param("id")

	oldNote := models.GetNoteById(id)

	if oldNote.AuthorID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	var reqNote NoteRequest
	c.BindJSON(&reqNote)

	var updatedNote models.Note

	updatedNote.ID = oldNote.ID
	updatedNote.AuthorID = oldNote.AuthorID
	updatedNote.Name = reqNote.Name
	updatedNote.Content = reqNote.Content

	resNote, err := updatedNote.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resNote)
}

func DeleteNote(c *gin.Context) {
	user, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}

	id := c.Param("id")
	note := models.DeleteNote(id)
	if note.AuthorID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
	}
	c.JSON(http.StatusOK, note)
}
