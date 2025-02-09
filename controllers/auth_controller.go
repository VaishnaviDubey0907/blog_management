package controllers

import (
	"blog-management/models"
	"blog-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := services.Login(user)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
