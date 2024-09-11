package ports_product

import domain_product "github.com/JoseLuis21/skeleton-lambda-api-go/internal/core/product/domain"

// Service ...
type Service interface {
	Find(code string) (*domain_product.Product, error)
	Store(product *domain_product.Product) error
	Update(product *domain_product.Product) error
	FindAll() ([]*domain_product.Product, error)
}

// Repository ...
type Repository interface {
	Find(code string) (*domain_product.Product, error)
	Store(product *domain_product.Product) error
	Update(product *domain_product.Product) error
	FindAll() ([]*domain_product.Product, error)
}
