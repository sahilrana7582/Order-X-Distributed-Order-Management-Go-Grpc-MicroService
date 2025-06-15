package handler

import (
	"context"

	pb "github.com/sahilrana7582/orderX/pkg/generated/product"
	"github.com/sahilrana7582/product-service/internal/domains"
	"github.com/sahilrana7582/product-service/internal/repository"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := &domains.Product{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Price:         req.GetPrice(),
		DiscountPrice: req.GetDiscountPrice(),
		Currency:      req.GetCurrency(),
		Status:        req.GetStatus().String(),
		Availability:  req.GetAvailability().String(),
		StockQuantity: req.GetStockQuantity(),
	}

	id, err := h.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Id:      id,
		Message: "Product created successfully",
	}, nil
}
