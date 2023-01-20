package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

type User struct {
	gorm.Model

	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     *string `gorm:"unique;not null" json:"email"`
	Avatar    string  `json:"avatar"`
	Password  string  `json:"-"`
	Entries   []Entry `json:"entries"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	EncryptedPassword := []byte(user.Password)

	if pw, err := bcrypt.GenerateFromPassword(EncryptedPassword, 0); err == nil {
		tx.Statement.SetColumn("password", string(pw))
	}

	return nil
}
