package controllers

import (
	"gin-learning/app"
	"gin-learning/middlewares"
	"gin-learning/services"
)

func CommentsController(a app.App) {
	sr := services.CommentsService{}
	sr.Models = a.Models
	{
		v1 := a.Router.Group("/comments")
		v1.Use(middlewares.Logger())
		v1.GET("", sr.AllComments)
	}
}
