package services

import (
	"blog-management/models"
	"fmt"
)

func Login(user models.User) string {
	fmt.Println("Logging in user:", user.Email)
	return "JWT-TOKEN"
}
