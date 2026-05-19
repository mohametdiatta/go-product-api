package controllers

import (
	"gin-learning/middlewares"
	"gin-learning/models"
	"gin-learning/mongorm"
	"gin-learning/services"
)

func CommentsController(a *mongorm.App) {
	comments, err := mongorm.GetModel[*models.Commentschema](a.Registry, "comments")
	if err != nil {
		panic(err)
	}
	{
		v1 := a.Router.Group("/comments")
		v1.Use(middlewares.Logger())
		v1.GET("", services.AllComments(comments))
	}
}
