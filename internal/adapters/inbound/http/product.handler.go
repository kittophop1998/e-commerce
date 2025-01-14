package http

import (
	"github.com/gofiber/fiber/v2"
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

func (p *ProductHandler) GetAll(c *fiber.Ctx) error {
	result := p.ProductService.GetAllProduct()

	return c.JSON(result)
}

func (p *ProductHandler) AddStock(c *fiber.Ctx) error {
	product := new(entities.Product)
	c.BodyParser(product)
	result := p.ProductService.AddStock(product)

	return c.JSON(result)
}
