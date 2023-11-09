package main

import (
	"github.com/Paulo-Prazeres/gocrud/initializers"
	"github.com/Paulo-Prazeres/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
