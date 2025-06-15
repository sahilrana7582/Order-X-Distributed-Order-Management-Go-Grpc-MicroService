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

func (r *productRepository) UpdateProduct(ctx context.Context, product *domains.Product) (string, error) {
	query := `
		UPDATE products SET
			name = $1,
			description = $2,
			price = $3,
			discount_price = $4,
			currency = $5,
			status = $6,
			availability = $7,
			stock_quantity = $8,
			updated_at = NOW()
		WHERE id = $9
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
		product.ID,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to update product: %w", err)
	}

	return id, nil
}
