package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oaraujocesar/perfinance/database"
	"github.com/oaraujocesar/perfinance/model"
)

func CreateEntry(c *gin.Context) {
	var input model.Entry

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})

		return
	}

	credit := model.Entry{Title: input.Title, Amount: input.Amount, TypeID: input.TypeID, CategoryID: input.CategoryID}
	database.DB.Create(&credit)

	c.JSON(http.StatusOK, gin.H{"data": credit})
}