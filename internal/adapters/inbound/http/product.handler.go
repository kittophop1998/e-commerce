package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kittiphop/e-commerce/internal/application/services"
	"github.com/kittiphop/e-commerce/internal/domain/entities"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (p *ProductHandler) GetAll(c *gin.Context) {
	result := p.ProductService.GetAllProduct()

	c.JSON(200, gin.H{"data": result})
}

func (p *ProductHandler) AddStock(c *gin.Context) {
	product := new(entities.Product)
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := p.ProductService.AddStock(product)

	c.JSON(200, gin.H{"data": result})
}
