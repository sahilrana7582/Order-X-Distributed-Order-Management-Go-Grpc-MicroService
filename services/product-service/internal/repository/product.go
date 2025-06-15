package repository

import (
	"context"

	"github.com/sahilrana7582/product-service/internal/domains"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domains.Product) (string, error)
	UpdateProduct(ctx context.Context, product *domains.Product) (string, error)
}
