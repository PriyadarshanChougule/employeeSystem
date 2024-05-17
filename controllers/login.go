package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type login struct {
	Uname    string `json:"uname" bson:"uname,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}

func VerifyLogin(c *gin.Context) {
	var loginDetails login
	err := c.BindJSON(&loginDetails)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	t, err := createToken(&loginDetails)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": t})
}

func createToken(loginDetails *login) (string, error) {
	var secretKey = "this is secret"
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userName": loginDetails.Uname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
