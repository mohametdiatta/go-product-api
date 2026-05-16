package main

import (
	"gin-learning/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	controllers.ProductController(app)
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Run(":3000")
}
