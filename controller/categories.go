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

func GetCategories(c *gin.Context) {
	var categories []model.Category

	result := database.DB.Find(&categories)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategory(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	result := database.DB.First(&category, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func UpdateCategory(c *gin.Context) {
	var input model.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})

		return
	}

	id := c.Param("id")

	result := database.DB.Model(&input).Where("id = ?", id).Update("name", input.Name)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No item with matching criteria was found on the database."})

		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error})

		return
	}

	database.DB.First(&input, id)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteCategory(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	result := database.DB.Delete(&category, id)
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
