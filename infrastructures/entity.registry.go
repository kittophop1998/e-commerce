package infrastructures

import "github.com/kittiphop/e-commerce/internal/domain/entities"

func RegisterEntities() []interface{} {
	return []interface{}{
		&entities.Product{},
	}
}
