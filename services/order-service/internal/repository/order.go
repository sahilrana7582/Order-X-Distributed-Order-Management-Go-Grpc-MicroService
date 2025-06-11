package repository

import (
	"context"

	"github.com/sahilrana7582/order-service/internal/domains"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domains.CreateOrderRequest) (*domains.CreateOrderResponse, error)
	GetAllOrdersRequest(ctx context.Context) (*domains.GetAllOrdersResponse, error)
	GetOrderByID(ctx context.Context, id string) (*domains.Order, error)
	UpdateOrder(ctx context.Context, updateOrder *domains.Order) error
}
