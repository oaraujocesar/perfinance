package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oaraujocesar/perfinance/database"
	"github.com/oaraujocesar/perfinance/model"
)

func CreateUser(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})

		return
	}

	user := model.User{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, Avatar: input.Avatar, Password: input.Password}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()})

		return
	}

	user.MarshalJSON()
	c.JSON(http.StatusOK, gin.H{"data": user})
}
