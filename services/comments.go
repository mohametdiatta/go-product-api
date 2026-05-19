package services

import (
	"gin-learning/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AllComments(comment *models.Commentschema) gin.HandlerFunc {
	return func(c *gin.Context) {

		var comments []models.Commentschema

		err := comment.FindAll(bson.D{}, &comments)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"count":   len(comments),
			"data":    comments,
		})
	}
}
