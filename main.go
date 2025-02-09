package main

import (
	"github.com/VaishnaviDubey0907/blog_management/config"
	"github.com/VaishnaviDubey0907/blog_management/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
