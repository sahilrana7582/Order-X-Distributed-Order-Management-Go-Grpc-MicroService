package app

import (
	"context"

	pb "github.com/sahilrana7582/orderX/pkg/generated/order"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:         "mock-order-id",
			UserId:     req.UserId,
			ProductId:  "mock-product-id",
			Quantity:   1,
			Status:     "PENDING",
			TotalPrice: "100.00",
		},
	}, nil
}
