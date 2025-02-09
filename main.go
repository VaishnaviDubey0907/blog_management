package main

import (
	"blog-management/config"
	routes "blog-management/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
