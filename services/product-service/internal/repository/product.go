package repository

import (
	"context"

	"github.com/sahilrana7582/product-service/internal/domains"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domains.Product) (string, error)
	UpdateProduct(ctx context.Context, product *domains.Product) (string, error)
	GetById(ctx context.Context, id string) (*domains.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetAll(ctx context.Context, page, limit int32) ([]*domains.Product, int32, error)
}
