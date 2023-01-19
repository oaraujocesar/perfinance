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

	router.GET("/entries", controller.GetEntries)
	router.GET("/entries/:id", controller.GetEntry)
	router.POST("/entries", controller.CreateEntry)

	router.GET("/categories", controller.GetCategories)
	router.GET("/categories/:id", controller.GetCategories)
	router.POST("/categories", controller.CreateCategory)
	router.PATCH("/categories/:id", controller.UpdateCategory)

	router.GET("/types", controller.GetTypes)
	router.GET("/types/:id", controller.GetType)
	router.POST("/types", controller.CreateType)

	router.Run()
}
