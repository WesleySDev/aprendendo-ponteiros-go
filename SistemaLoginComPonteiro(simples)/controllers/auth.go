package controllers

import (
	"login-go/database"
	"login-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados iválidos"})
		return
	}

	database.USers = append(database.USers, user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário cadastrado",
	})
}

func Login(c *gin.Context) {
	var input models.User

	c.ShouldBindJSON(&input)
	for _, user := range database.Users {

		if user.Email == input.Email && user.Password == input.password {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login realizado",
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Credenciais inválidas",
	})
}
