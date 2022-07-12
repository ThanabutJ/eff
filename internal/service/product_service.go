package service

import (
	"context"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductService interface {
	FindAll(ctx context.Context) ([]*Product, error)
}

type productService struct {
	productRespository ProductRespository
}

func NewProductService(productRepo ProductRespository) ProductService {
	return &productService{
		productRespository: productRepo,
	}
}

func (s *productService) FindAll(ctx context.Context) ([]*Product, error) {
	products, err := s.productRespository.FindAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

type ProductRespository interface {
	FindAllProduct(ctx context.Context) ([]*Product, error)
}
