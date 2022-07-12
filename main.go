package main

import (
	"github.com/ThanabutJ/eff/internal/app"
	"github.com/ThanabutJ/eff/internal/repositories"
	"github.com/ThanabutJ/eff/internal/repositories/productserviceadapter"
	"github.com/ThanabutJ/eff/internal/server"
	"github.com/ThanabutJ/eff/internal/service"
)

func main() {
	productRepository := repositories.NewProductRepository()
	productRepoAdaptor := productserviceadapter.NewProductRepositoryAdapter(productRepository)
	productService := service.NewProductService(productRepoAdaptor)
	app := app.New(productService)

	handlers := server.NewHandlers(app)
	server := server.New(handlers)

	server.Start()
}
