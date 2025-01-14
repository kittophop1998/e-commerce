package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittiphop/e-commerce/infrastructures"
	"github.com/kittiphop/e-commerce/infrastructures/db/gorm"
	"github.com/kittiphop/e-commerce/internal/adapters/inbound/http"
	"github.com/kittiphop/e-commerce/internal/application/services"
	"github.com/kittiphop/e-commerce/internal/application/usecases"
	"github.com/kittiphop/e-commerce/internal/routers"
)

func InitializeApp() (*fiber.App, error) {
	// Setup Database
	db := infrastructures.NewDatabase()
	infrastructures.AutoMigrate(db)

	// Setup Product
	productRepo := gorm.NewProductRepo(db)
	productUseCase := usecases.NewManageStockUseCase(productRepo)
	productService := services.NewProductService(*productUseCase)
	productHandler := http.NewProductHandler(*productService)

	// Setup Fiber Router
	app := fiber.New()
	routers.ProductSetupRouter(app, productHandler)

	return app, nil
}
