package app

import (
	"context"
	"log"

	"github.com/sahilrana7582/order-service/internal/domains"
	"github.com/sahilrana7582/order-service/internal/repository"
	pb "github.com/sahilrana7582/orderX/pkg/generated/order"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	OrderRepo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{
		OrderRepo: repo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderReq := &domains.CreateOrderRequest{
		CustomerID:     req.CustomerId,
		OrderNumber:    req.OrderNumber,
		OrderStatus:    domains.OrderStatus(req.OrderStatus.String()),
		PaymentStatus:  domains.PaymentStatus(req.PaymentStatus.String()),
		ShippingMethod: domains.ShippingMethod(req.ShippingMethod.String()),
		TotalAmount:    req.TotalAmount,
		Currency:       req.Currency,
	}

	res, err := s.OrderRepo.CreateOrder(ctx, orderReq)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:      res.ID,
		Message: res.Message,
	}, nil
}
