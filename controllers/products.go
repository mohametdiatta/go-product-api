package controllers

import (
	"gin-learning/app"
	"gin-learning/middlewares"
	"gin-learning/services"
)

func ProductController(a app.App) {
	sr := services.ProductsService{}
	sr.Models = a.Models
	{
		v1 := a.Router.Group("/products")
		v1.Use(middlewares.Logger())
		v1.POST("", sr.CreateProduct)
		v1.GET("", sr.AllProducts)
		v1.DELETE("/:id", sr.DeleteProduct)
	}
}
