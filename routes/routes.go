package routes

import (
	"github.com/VaishnaviDubey0907/blog_management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/blog", controllers.CreateBlog)
	r.GET("/blog/:id", controllers.GetBlogByID)

	return r
}
