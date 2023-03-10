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

func GetEntries(c *gin.Context) {
	var entries []model.Entry

	result := database.DB.Find(&entries)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": entries})
}

func GetEntry(c *gin.Context) {
	var entry model.Entry

	id := c.Param("id")

	result := database.DB.First(&entry, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": entry})
}

func UpdateEntry(c *gin.Context) {
	var input model.Entry

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	id := c.Param("id")
	result := database.DB.Model(&input).Where("id = ?", id).Updates(input)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No item with matching criteria was found on the database."})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	database.DB.First(&input, id)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteEntry(c *gin.Context) {
	var entry model.Entry

	id := c.Param("id")

	result := database.DB.Delete(&entry, id)
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
