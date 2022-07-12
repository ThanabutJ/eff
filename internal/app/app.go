package app

import "github.com/ThanabutJ/eff/internal/service"

type Appllication struct {
	ProductService service.ProductService
}

func New(productService service.ProductService) *Appllication {
	return &Appllication{
		ProductService: productService,
	}
}
