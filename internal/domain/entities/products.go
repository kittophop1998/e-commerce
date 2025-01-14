package entities

type Product struct {
	ID           int     `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	CreationDate string  `json:"creationDate"`
	LastUpdate   string  `json:"lastUpdate"`
}

func (p *Product) Update(product *Product) {
	p.Name = product.Name
	p.Description = product.Description
	p.Price = product.Price
	p.LastUpdate = product.LastUpdate
}

func (p *Product) InValidProduct(product *Product) bool {
	return product.Name == "" || product.Description == "" || product.Price == 0
}
