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

func (r *productRepository) GetById(ctx context.Context, id string) (*domains.Product, error) {
	query := `
		SELECT
			id, name, description, price, discount_price,
			currency, status, availability, stock_quantity,
			created_at, updated_at
		FROM products
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var product domains.Product

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.DiscountPrice,
		&product.Currency,
		&product.Status,
		&product.Availability,
		&product.StockQuantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to fetch product by ID: %w", err)
	}

	return &product, nil

}

func (r *productRepository) DeleteProduct(ctx context.Context, id string) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *productRepository) GetAll(ctx context.Context, page, limit int32) ([]*domains.Product, int32, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := `
		SELECT * from products
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []*domains.Product

	for rows.Next() {
		var p domains.Product

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.DiscountPrice,
			&p.Currency,
			&p.Status,
			&p.Availability,
			&p.StockQuantity,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &p)
	}

	var total int32
	countQuery := `SELECT COUNT(*) FROM products`
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	return products, total, nil

}
