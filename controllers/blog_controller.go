package controllers

import (
	"blog-management/models"
	"blog-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.CreateBlog(blog)
	c.JSON(http.StatusOK, gin.H{"message": "Blog created"})
}
