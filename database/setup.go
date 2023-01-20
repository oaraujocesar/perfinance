package database

import (
	"fmt"
	"os"

	m "github.com/oaraujocesar/perfinance/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=America/Sao_Paulo\n", host, user, password, name, port)

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("DATABASE: Failed to connect.")
	}

	err = database.AutoMigrate(&m.User{}, &m.Type{}, &m.Entry{}, &m.Category{})
	if err != nil {
		return
	}

	DB = database
}
