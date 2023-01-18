package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oaraujocesar/perfinance/controller"
	"github.com/oaraujocesar/perfinance/database"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/entries", controller.CreateEntry)

	router.POST("/categories", controller.CreateCategory)

	router.POST("/types", controller.CreateType)
	router.GET("/types", controller.GetTypes)
	router.GET("/types/:id", controller.GetType)

	router.Run()
}
