package repository

import (
	"context"

	"github.com/sahilrana7582/order-service/internal/domains"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domains.CreateOrderRequest) (*domains.CreateOrderResponse, error)
	GetAllOrdersRequest(ctx context.Context) (*domains.GetAllOrdersResponse, error)
}
