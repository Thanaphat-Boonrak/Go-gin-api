package routes

import (
	"net/http"

	"Go-gin-rest-api/models"
	"Go-gin-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot bind data",
			"error":   err.Error(),
		})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot save user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot bind data",
			"error":   err.Error(),
		})
		return
	}

	if err := user.ValidateCredential(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
