package repositories

import (
	"context"
	"time"
)

// ProductInDB
// Why there are more than 1 Product struct.
// read https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
type ProductInDB struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*ProductInDB, error)
}

// productRepository
type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

// FindAll
// We suppose to do actual DB query here.
// However our struct dont have any DB connection setup yet. So I will just left if like this.
func (r *productRepository) FindAll(ctx context.Context) ([]*ProductInDB, error) {
	return []*ProductInDB{}, nil
}
