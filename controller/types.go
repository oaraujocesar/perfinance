package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oaraujocesar/perfinance/database"
	"github.com/oaraujocesar/perfinance/model"
)

func CreateType(c *gin.Context) {
	var input model.Type

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})

		return
	}

	entryType := model.Type{Name: input.Name}
	database.DB.Create(&entryType)

	c.JSON(http.StatusOK, gin.H{"data": entryType})
}
