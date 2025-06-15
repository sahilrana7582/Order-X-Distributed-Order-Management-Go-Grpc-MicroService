package handler

import (
	"context"
	"database/sql"
	"errors"

	pb "github.com/sahilrana7582/orderX/pkg/generated/product"
	"github.com/sahilrana7582/product-service/internal/domains"
	"github.com/sahilrana7582/product-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product := &domains.Product{
		ID:            req.GetId(),
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Price:         req.GetPrice(),
		DiscountPrice: req.GetDiscountPrice(),
		Currency:      req.GetCurrency(),
		Status:        req.GetStatus().String(),
		Availability:  req.GetAvailability().String(),
		StockQuantity: req.GetStockQuantity(),
	}

	id, err := h.repo.UpdateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProductResponse{
		Id:      id,
		Message: "Product updated successfully",
	}, nil
}

func (h *ProductHandler) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	product, err := h.repo.GetById(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &pb.GetProductByIdResponse{
		Product: &pb.Product{
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Currency:      product.Currency,
			Status:        pb.ProductStatus(pb.ProductStatus_value[product.Status]),
			Availability:  pb.AvailabilityStatus(pb.AvailabilityStatus_value[product.Availability]),
			StockQuantity: int32(product.StockQuantity),
			CreatedAt:     product.CreatedAt.String(),
			UpdatedAt:     product.UpdatedAt.String(),
		},
	}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := h.repo.DeleteProduct(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}

	return &pb.DeleteProductResponse{
		Message: "Product deleted successfully",
	}, nil
}
