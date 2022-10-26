package main

import (
	"github.com/andsholinka/golang-gin/controllers"
	"github.com/andsholinka/golang-gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.CoonetToDB()
}

// var (
// 	DB *gorm.DB = initializers.SetupDBConnection()
// )

func main() {
	// defer initializers.CloseDBConnection(DB)
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PATCH("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
