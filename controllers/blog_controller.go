package controllers

import (
	"net/http"

	"github.com/VaishnaviDubey0907/blog_management/models"
	"github.com/VaishnaviDubey0907/blog_management/services"

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
func GetBlogByID(c *gin.Context) {
	id := c.Param("id") // Get the blog ID from URL
	blog, err := services.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}
