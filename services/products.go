package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsService struct {
	c *gin.Context
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
