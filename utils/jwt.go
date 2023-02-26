package utils

import (
	"net/http"
	"os"

	j "github.com/cristalhq/jwt/v4"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateJWT(c *gin.Context, sub string) (jwt string) {
	port := os.Getenv("PORT")
	secret := os.Getenv("JWT_SECRET")
	key := []byte(secret)

	signer, err := j.NewSignerHS(j.HS256, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error})

		return
	}

	claims := &j.RegisteredClaims{
		Audience: []string{"http://localhost:" + port},
		ID:       uuid.New().String(),
		Subject:  sub,
	}

	builder := j.NewBuilder(signer)

	token, err := builder.Build(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error})

		return
	}

	return token.String()
}
