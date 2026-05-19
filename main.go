package main

import (
	"context"
	"gin-learning/controllers"
	"gin-learning/models"
	"gin-learning/mongorm"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGO_DB_URL")
	client, err := mongorm.Connect(url)
	if err != nil {
		panic(err)
	}

	db := client.Database("sample_mflix")
	registry := mongorm.NewRegistry(context.Background(), db)
	registry.Register("comments", &models.Commentschema{}, "comments")

	app := mongorm.NewApp(registry)

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	controllers.CommentsController(app)
	app.Run(":3000")
}
