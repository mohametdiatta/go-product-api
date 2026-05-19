package controllers

import (
	"gin-learning/middlewares"
	"gin-learning/models"

	"gin-learning/services"

	"github.com/mohametdiatta/gormongo"
)

func CommentsController(a *gormongo.App) {
	comments, err := gormongo.GetModel[*models.Commentschema](a.Registry, "comments")
	if err != nil {
		panic(err)
	}
	{
		v1 := a.Router.Group("/comments")
		v1.Use(middlewares.Logger())
		v1.GET("", services.AllComments(comments))
	}
}
