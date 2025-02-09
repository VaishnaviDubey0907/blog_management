package controllers

import (
	"net/http"

	"github.com/VaishnaviDubey0907/blog_management/models"
	"github.com/VaishnaviDubey0907/blog_management/services"
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
