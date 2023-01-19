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

func GetTypes(c *gin.Context) {
	var types []model.Type

	result := database.DB.Find(&types)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": types})
}

func GetType(c *gin.Context) {
	var entityType model.Type

	id := c.Param("id")

	result := database.DB.First(&entityType, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": entityType})
}

func DeleteType(c *gin.Context) {
	var entityType model.Type

	id := c.Param("id")

	result := database.DB.Delete(&entityType, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No item with matching criteria was found on the database."})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	c.Status(http.StatusOK)
}
