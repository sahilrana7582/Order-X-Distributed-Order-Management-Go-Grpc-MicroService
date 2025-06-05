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

func (s *OrderService) GetAllOrders(ctx context.Context, req *pb.GetAllOrdersRequest) (*pb.GetAllOrdersResponse, error) {
	res, err := s.OrderRepo.GetAllOrdersRequest(ctx)
	if err != nil {
		log.Printf("Error fetching all orders: %v", err)
		return nil, err
	}

	var orders []*pb.Order
	for _, order := range res.Orders {
		orders = append(orders, &pb.Order{
			Id:             order.ID,
			CustomerId:     order.CustomerID,
			OrderNumber:    order.OrderNumber,
			OrderStatus:    pb.OrderStatus(pb.OrderStatus_value[string(order.OrderStatus)]),
			PaymentStatus:  pb.PaymentStatus(pb.PaymentStatus_value[string(order.PaymentStatus)]),
			ShippingMethod: pb.ShippingMethod(pb.ShippingMethod_value[string(order.ShippingMethod)]),
			TotalAmount:    order.TotalAmount,
			Currency:       order.Currency,
		})
	}

	return &pb.GetAllOrdersResponse{
		Orders: orders,
	}, nil
}
