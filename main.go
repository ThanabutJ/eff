package main

import (
	"eff/internal/app"
	"eff/internal/repositories"
	"eff/internal/repositories/productserviceadapter"
	"eff/internal/server"
	"eff/internal/service"
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
