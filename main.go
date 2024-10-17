package main

import (
	"learnApi/controllers"
	"learnApi/middleware"
	"learnApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectToDatabase()

	router.Use(middleware.RequestResponseLogger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/post", controllers.StorePost)
	router.GET("/api/post/:id", controllers.FindPostById)
	router.PUT("/api/post/:id", controllers.UpdatePost)
	router.DELETE("/api/post/:id", controllers.DeletePost)

	router.Run(":8000")
}
