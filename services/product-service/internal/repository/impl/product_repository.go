package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sahilrana7582/product-service/internal/domains"
	"github.com/sahilrana7582/product-service/internal/repository"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *domains.Product) (string, error) {
	query := `
		INSERT INTO products (
			name, description, price, discount_price,
			currency, status, availability, stock_quantity,
			created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW())
		RETURNING id
	`

	var id string
	err := r.db.QueryRowContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.DiscountPrice,
		product.Currency,
		product.Status,
		product.Availability,
		product.StockQuantity,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to create product: %w", err)
	}

	return id, nil
}
