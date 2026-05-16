package controllers

import (
	"gin-learning/services"

	"github.com/gin-gonic/gin"
)

func ProductController(app *gin.Engine) {
	sr := services.ProductsService{}
	{
		v1 := app.Group("/products")
		v1.POST("", sr.CreateProduct)
		v1.GET("", sr.AllProducts)
		v1.DELETE("/:id", sr.DeleteProduct)
	}
}
