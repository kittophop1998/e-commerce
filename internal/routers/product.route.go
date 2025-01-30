package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kittiphop/e-commerce/internal/adapters/inbound/http"
)

func ProductSetupRouter(app *gin.Engine, handler *http.ProductHandler) {
	productApi := app.Group("/api/products")
	productApi.POST("/add-stock", handler.AddStock)
	productApi.GET("/get-stock-all", handler.GetAll)
	// productApi.Get("/products/:id", handler.GetByID)
	// productApi.Put("/products/:id", handler.Update)
	// productApi.Delete("/products/:id", handler.Delete)
}
