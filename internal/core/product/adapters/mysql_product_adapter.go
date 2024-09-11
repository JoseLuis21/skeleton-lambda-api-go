package adapters_product

import (
	domain_product "github.com/JoseLuis21/skeleton-lambda-api-go/internal/core/product/domain"
	ports_product "github.com/JoseLuis21/skeleton-lambda-api-go/internal/core/product/ports"
)

var _ ports_product.Repository = (*ProductRepository)(nil)

type ProductRepository struct {
	repo any
}

// NewProductRepository ...
func NewProductRepository(repo any) *ProductRepository {
	return &ProductRepository{
		repo: repo,
	}
}

func (r *ProductRepository) Find(code string) (*domain_product.Product, error) {
	panic("implement me")
}

func (r *ProductRepository) Store(product *domain_product.Product) error {
	panic("implement me")
}

func (r *ProductRepository) Update(product *domain_product.Product) error {
	panic("implement me")
}

func (r *ProductRepository) FindAll() ([]*domain_product.Product, error) {
	panic("implement me")
}
