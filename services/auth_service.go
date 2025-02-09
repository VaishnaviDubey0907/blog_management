package services

import (
	"fmt"

	"github.com/VaishnaviDubey0907/blog_management/models"
)

func Login(user models.User) string {
	fmt.Println("Logging in user:", user.Email)
	return "JWT-TOKEN"
}
