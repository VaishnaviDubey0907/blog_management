package services

import (
	"blog-management/models"
	"fmt"
)

func CreateBlog(blog models.Blog) {
	fmt.Println("Creating blog:", blog.Title)
}
