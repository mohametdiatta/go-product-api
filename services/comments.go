package services

import (
	"gin-learning/database"
	"gin-learning/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentsService struct {
	Models database.ModelRegistery
}

func (p *CommentsService) AllComments(c *gin.Context) {
	// data:=p.Models.Model["comments"].FindAll()
	// c.JSON(http.StatusOK, gin.H{"data": "List AllProducts"})

	var comments []models.Commentschema

	err := p.Models.Model["comments"].FindAll(bson.D{}, &comments) // bson.D{} = pas de filtre
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}
