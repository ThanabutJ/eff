package server

import (
	"eff/internal/app"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetAllProducts(c *gin.Context)
}

type handlers struct {
	app *app.Appllication
}

func NewHandlers(app *app.Appllication) Handlers {
	return &handlers{
		app: app,
	}
}

func (h *handlers) GetAllProducts(c *gin.Context) {
	products, err := h.app.ProductService.FindAll(c)
	if err != nil {
		c.JSON(500, gin.H{"status": "Internal Server Error", "error": err})
		return
	}

	c.JSON(200, gin.H{"status": "OK", "data": products})
}
