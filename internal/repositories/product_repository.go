package repositories

import "context"

type ProductInDB struct {
	ID   string
	Name string
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*ProductInDB, error)
}

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll(ctx context.Context) ([]*ProductInDB, error) {
	return []*ProductInDB{}, nil
}
