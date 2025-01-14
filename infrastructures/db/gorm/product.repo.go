package gorm

import (
	"github.com/kittiphop/e-commerce/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// Implement function here
func (p *ProductRepo) GetAll() *[]entities.Product {
	var products []entities.Product
	p.db.Find(&products)

	return &products
}

func (p *ProductRepo) Create(product *entities.Product) (*entities.Product, error) {
	result := p.db.Create(&product)

	return product, result.Error
}
