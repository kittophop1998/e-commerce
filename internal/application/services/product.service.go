package services

import (
	"github.com/kittiphop/e-commerce/internal/application/usecases"
	"github.com/kittiphop/e-commerce/internal/domain/entities"
)

type ProductService struct {
	ManageStockUseCase usecases.ManageStockUseCase
}

func NewProductService(manageStockUseCase usecases.ManageStockUseCase) *ProductService {
	return &ProductService{ManageStockUseCase: manageStockUseCase}
}

func (p *ProductService) GetAllProduct() *[]entities.Product {
	return p.ManageStockUseCase.GetAllProduct()
}

func (p *ProductService) AddStock(product *entities.Product) *entities.Product {
	return p.ManageStockUseCase.AddStock(product)
}
