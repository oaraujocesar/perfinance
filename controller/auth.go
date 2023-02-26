package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"github.com/oaraujocesar/perfinance/database"
	"github.com/oaraujocesar/perfinance/model"
	"github.com/oaraujocesar/perfinance/utils"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signin(c *gin.Context) {
	var auth Auth
	var user model.User

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})

		return
	}

	result := database.DB.Where("email = ?", auth.Email).First(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Usuário não cadastrado."})

		return
	}
	fmt.Println(auth.Password, user.Password)
	ok, err := argon2.VerifyEncoded([]byte(auth.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})

		return
	}

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Senha incorreta."})

		return
	}

	// se existir, crio o JWT
	token := utils.CreateJWT(c, *user.Email)

	// envio o token como resposta

	c.JSON(http.StatusOK, gin.H{"data": token})
}
