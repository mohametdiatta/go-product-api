package services

import (
	"gin-learning/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsService struct {
	Models database.ModelRegistery
}

func (p *ProductsService) CreateProduct(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"action": "Create Products"})

}
func (p *ProductsService) AllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "List AllProducts"})

}
func (p *ProductsService) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"action": "DeleteProduct", "id": id})

}
