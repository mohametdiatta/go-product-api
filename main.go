package main

import (
	"context"
	"gin-learning/app"
	"gin-learning/controllers"
	"gin-learning/database"
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
	registery := database.ModelRegistery{}
	registery.New(context.Background(), db)
	registery.Register(map[string]*mongorm.Model{
		"comments": models.Comment,
	}, "comments")
	registery.Init()

	app := &app.App{
		Router: gin.Default(),
		Models: registery,
	}

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	controllers.ProductController(*app)
	controllers.CommentsController(*app)
	app.Run()
}
