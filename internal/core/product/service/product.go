package service_product

import (
	domain_product "github.com/JoseLuis21/skeleton-lambda-api-go/internal/core/product/domain"
	ports_product "github.com/JoseLuis21/skeleton-lambda-api-go/internal/core/product/ports"
)

var _ ports_product.Service = (*ProductService)(nil)

type ProductService struct {
	productrepo ports_product.Repository
}

// NewProductService ...
func NewProductService(productrepo ports_product.Repository) *ProductService {
	return &ProductService{productrepo: productrepo}
}

func (s *ProductService) Find(code string) (*domain_product.Product, error) {
	return s.productrepo.Find(code)
}

func (s *ProductService) Store(product *domain_product.Product) error {
	return s.productrepo.Store(product)

}
func (s *ProductService) Update(product *domain_product.Product) error {
	return s.productrepo.Update(product)
}

func (s *ProductService) FindAll() ([]*domain_product.Product, error) {
	return s.productrepo.FindAll()
}
