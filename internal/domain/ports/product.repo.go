package ports

import "github.com/kittiphop/e-commerce/internal/domain/entities"

type ProductRepo interface {
	GetAll() *[]entities.Product
	Create(product *entities.Product) (*entities.Product, error)
}
