package main

import (
	"github.com/andsholinka/golang-gin/initializers"
	"github.com/andsholinka/golang-gin/models"
)

func init() {
	// initializers.LoadEnvVariables()
	initializers.CoonetToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
