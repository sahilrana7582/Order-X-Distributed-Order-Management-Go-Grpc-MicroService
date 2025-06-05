package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sahilrana7582/order-service/internal/domains"
	"github.com/sahilrana7582/order-service/internal/repository"
)

type OrderRepositoryImpl struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &OrderRepositoryImpl{
		DB: db,
	}
}

func (repo *OrderRepositoryImpl) CreateOrder(ctx context.Context, order *domains.CreateOrderRequest) (*domains.CreateOrderResponse, error) {
	query := `
		INSERT INTO orders (
			customer_id,
			order_number,
			order_status,
			payment_status,
			shipping_method,
			total_amount,
			currency
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	var id string

	err := repo.DB.QueryRowContext(ctx, query,
		order.CustomerID,
		order.OrderNumber,
		order.OrderStatus,
		order.PaymentStatus,
		order.ShippingMethod,
		order.TotalAmount,
		order.Currency,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	return &domains.CreateOrderResponse{
		ID:      id,
		Message: "Order created successfully",
	}, nil
}

func (repo *OrderRepositoryImpl) GetAllOrdersRequest(ctx context.Context) (*domains.GetAllOrdersResponse, error) {
	query := `
		SELECT id, customer_id, order_number, order_status, payment_status, shipping_method, total_amount, currency
		FROM orders;
	`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	var orders []domains.Order
	for rows.Next() {
		var order domains.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.OrderNumber,
			&order.OrderStatus, &order.PaymentStatus,
			&order.ShippingMethod, &order.TotalAmount, &order.Currency); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}

	return &domains.GetAllOrdersResponse{Orders: orders}, nil
}
