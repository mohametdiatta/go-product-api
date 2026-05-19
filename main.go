package main

import (
	"context"
	"gin-learning/controllers"
	"gin-learning/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohametdiatta/gormongo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGO_DB_URL")
	client, err := gormongo.Connect(url)
	if err != nil {
		panic(err)
	}

	db := client.Database("sample_mflix")
	registry := gormongo.NewRegistry(context.Background(), db)
	registry.Register("comments", &models.Commentschema{}, "comments")

	app := gormongo.NewApp(registry)

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	controllers.CommentsController(app)
	app.Run(":3000")
}
