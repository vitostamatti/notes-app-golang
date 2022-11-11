package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Password    string `gorm:"size:255;not null;" json:"password"`
	IsSuperuser bool   `gorm:"default:false" json:"is_superuser"`
	Notes       []Note `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"notes"`
}

func (u *User) Create() (*User, error) {

	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return &User{}, err
	}

	u.Password = string(hashedPassword)

	err = DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) Save() (*User, error) {

	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return &User{}, err
	}

	u.Password = string(hashedPassword)

	err = DB.Save(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func GetUsers() []User {
	var Users []User
	DB.Find(&Users)
	return Users
}

func GetUserById(ID string) (User, error) {
	var user User
	err := DB.Where("ID=?", ID).Find(&user).Error
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := DB.Where("username=?", username).Find(&user).Error
	return user, err
}

func DeleteUser(ID string) (User, error) {
	var user User
	err := DB.Where("ID=?", ID).Delete(&user).Error
	return user, err
}
