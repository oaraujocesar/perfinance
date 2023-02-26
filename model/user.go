package model

import (
	"encoding/json"
	"fmt"

	"github.com/matthewhartstonge/argon2"

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
	Password  string  `json:"password"`
	Entries   []Entry `json:"entries"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	argon := argon2.DefaultConfig()
	hashedPassword, err := argon.HashEncoded([]byte(user.Password))
	fmt.Println(user.Password, string(hashedPassword))
	if err == nil {
		tx.Statement.SetColumn("password", string(hashedPassword))
	}

	return nil
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	safeUser := struct {
		Password string `json:"password,omitempty"`
		Alias
	}{
		Alias: Alias(u),
	}

	return json.Marshal(safeUser)
}
