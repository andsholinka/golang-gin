package main

import (
	"github.com/andsholinka/golang-gin/controllers"
	"github.com/andsholinka/golang-gin/initializers"
	"github.com/andsholinka/golang-gin/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.CoonetToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
