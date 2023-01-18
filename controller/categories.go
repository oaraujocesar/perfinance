package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oaraujocesar/perfinance/database"
	"github.com/oaraujocesar/perfinance/model"
)

func CreateCategory(c *gin.Context) {
	var input model.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})

		return
	}

	category := model.Category{Name: input.Name}
	database.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}
