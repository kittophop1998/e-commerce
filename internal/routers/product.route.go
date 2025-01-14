package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittiphop/e-commerce/internal/adapters/inbound/http"
)

func ProductSetupRouter(app *fiber.App, handler *http.ProductHandler) {
	productApi := app.Group("/api/products")
	productApi.Post("/add-stock", handler.AddStock)
	productApi.Get("/get-stock-all", handler.GetAll)
	// productApi.Get("/products/:id", handler.GetByID)
	// productApi.Put("/products/:id", handler.Update)
	// productApi.Delete("/products/:id", handler.Delete)
}
