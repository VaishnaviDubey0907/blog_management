package services

import (
	"context"
	"fmt"

	"github.com/VaishnaviDubey0907/blog_management/config"
	"github.com/VaishnaviDubey0907/blog_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBlog(blog models.Blog) {
	fmt.Println("Creating blog:", blog.Title)
}
func GetBlogByID(id string) (models.Blog, error) {
	collection := config.DB.Database("blogdb").Collection("blogs") // Adjust DB name
	var blog models.Blog

	objID, err := primitive.ObjectIDFromHex(id) // Convert string ID to MongoDB ObjectID
	if err != nil {
		return blog, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&blog)
	if err != nil {
		return blog, err
	}

	return blog, nil
}
