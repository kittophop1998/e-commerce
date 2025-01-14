package usecases

import (
	"log"

	"github.com/kittiphop/e-commerce/internal/domain/entities"
	"github.com/kittiphop/e-commerce/internal/domain/ports"
)

type ManageStockUseCase struct {
	ProductRepo ports.ProductRepo
}

func NewManageStockUseCase(productRepo ports.ProductRepo) *ManageStockUseCase {
	return &ManageStockUseCase{ProductRepo: productRepo}
}

func (m *ManageStockUseCase) AddStock(product *entities.Product) *entities.Product {
	result, error := m.ProductRepo.Create(product)
	if error != nil {
		log.Fatalf("failed to add stock: %v", error)
	}

	return result
}

func (m *ManageStockUseCase) GetAllProduct() *[]entities.Product {
	products := m.ProductRepo.GetAll()

	return products
}
