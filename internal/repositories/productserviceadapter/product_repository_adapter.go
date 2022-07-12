package productserviceadapter

import (
	"context"

	"github.com/ThanabutJ/eff/internal/repositories"
	"github.com/ThanabutJ/eff/internal/service"
)

type productRepositoryAdapter struct {
	repo repositories.ProductRepository
}

func NewProductRepositoryAdapter(repo repositories.ProductRepository) service.ProductRespository {
	return &productRepositoryAdapter{
		repo: repo,
	}
}

func (a *productRepositoryAdapter) FindAllProduct(ctx context.Context) ([]*service.Product, error) {
	dbProducts, err := a.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]*service.Product, len(dbProducts))
	for i, p := range dbProducts {
		products[i] = mapDBProductToServiceProduct(p)
	}

	return products, nil
}

func mapDBProductToServiceProduct(p *repositories.ProductInDB) *service.Product {
	return &service.Product{
		ID:   p.ID,
		Name: p.Name,
	}
}
