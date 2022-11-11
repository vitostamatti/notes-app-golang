package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createSuperUser() {
	username := os.Getenv("SUPERUSER_USERNAME")
	password := os.Getenv("SUPERUSER_PASSWORD")
	user := User{
		Username:    username,
		Password:    password,
		IsSuperuser: true,
	}
	user.Create()
}

func ConnectDataBase() {

	dbSchema := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":5432/" + dbSchema
	// dbURL := "postgres://admin:admin@0.0.0.0:5432/books"

	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d

	DB.AutoMigrate(&User{}, &Note{})

	createSuperUser()

}
