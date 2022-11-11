package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Name     string `json:"name"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
}

func (n *Note) Create() (*Note, error) {
	result := DB.Create(&n)
	if result.Error != nil {
		return &Note{}, result.Error
	}

	return n, nil
}
func (n *Note) Save() (*Note, error) {
	result := DB.Save(&n)
	if result.Error != nil {
		return &Note{}, result.Error
	}

	return n, nil
}

func GetNotes() []Note {
	var notes []Note
	DB.Find(&notes)
	return notes
}

func GetNoteById(ID string) Note {
	var note Note
	DB.Where("ID=?", ID).Find(&note)
	return note
}

func GetNotesByAuthor(author_id string) []Note {
	var notes []Note
	DB.Where("author_id=?", author_id).Find(&notes)
	return notes
}

func DeleteNote(ID string) Note {
	var note Note
	DB.Where("ID=?", ID).Delete(&note)
	return note
}
