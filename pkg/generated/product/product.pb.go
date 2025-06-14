// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: product/product.proto

package product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductStatus int32

const (
	ProductStatus_ACTIVE       ProductStatus = 0
	ProductStatus_INACTIVE     ProductStatus = 1
	ProductStatus_OUT_OF_STOCK ProductStatus = 2
	ProductStatus_DISCONTINUED ProductStatus = 3
)

// Enum value maps for ProductStatus.
var (
	ProductStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
		2: "OUT_OF_STOCK",
		3: "DISCONTINUED",
	}
	ProductStatus_value = map[string]int32{
		"ACTIVE":       0,
		"INACTIVE":     1,
		"OUT_OF_STOCK": 2,
		"DISCONTINUED": 3,
	}
)

func (x ProductStatus) Enum() *ProductStatus {
	p := new(ProductStatus)
	*p = x
	return p
}

func (x ProductStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_product_product_proto_enumTypes[0].Descriptor()
}

func (ProductStatus) Type() protoreflect.EnumType {
	return &file_product_product_proto_enumTypes[0]
}

func (x ProductStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductStatus.Descriptor instead.
func (ProductStatus) EnumDescriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

type AvailabilityStatus int32

const (
	AvailabilityStatus_AVAILABLE AvailabilityStatus = 0
	AvailabilityStatus_BACKORDER AvailabilityStatus = 1
	AvailabilityStatus_PREORDER  AvailabilityStatus = 2
)

// Enum value maps for AvailabilityStatus.
var (
	AvailabilityStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "BACKORDER",
		2: "PREORDER",
	}
	AvailabilityStatus_value = map[string]int32{
		"AVAILABLE": 0,
		"BACKORDER": 1,
		"PREORDER":  2,
	}
)

func (x AvailabilityStatus) Enum() *AvailabilityStatus {
	p := new(AvailabilityStatus)
	*p = x
	return p
}

func (x AvailabilityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvailabilityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_product_product_proto_enumTypes[1].Descriptor()
}

func (AvailabilityStatus) Type() protoreflect.EnumType {
	return &file_product_product_proto_enumTypes[1]
}

func (x AvailabilityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvailabilityStatus.Descriptor instead.
func (AvailabilityStatus) EnumDescriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice float64                `protobuf:"fixed64,6,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Currency      string                 `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        ProductStatus          `protobuf:"varint,8,opt,name=status,proto3,enum=product.ProductStatus" json:"status,omitempty"`
	Availability  AvailabilityStatus     `protobuf:"varint,9,opt,name=availability,proto3,enum=product.AvailabilityStatus" json:"availability,omitempty"`
	StockQuantity int32                  `protobuf:"varint,10,opt,name=stock_quantity,json=stockQuantity,proto3" json:"stock_quantity,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *Product) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Product) GetStatus() ProductStatus {
	if x != nil {
		return x.Status
	}
	return ProductStatus_ACTIVE
}

func (x *Product) GetAvailability() AvailabilityStatus {
	if x != nil {
		return x.Availability
	}
	return AvailabilityStatus_AVAILABLE
}

func (x *Product) GetStockQuantity() int32 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

func (x *Product) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice float64                `protobuf:"fixed64,5,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Currency      string                 `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        ProductStatus          `protobuf:"varint,8,opt,name=status,proto3,enum=product.ProductStatus" json:"status,omitempty"`
	Availability  AvailabilityStatus     `protobuf:"varint,9,opt,name=availability,proto3,enum=product.AvailabilityStatus" json:"availability,omitempty"`
	StockQuantity int32                  `protobuf:"varint,10,opt,name=stock_quantity,json=stockQuantity,proto3" json:"stock_quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_product_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *CreateProductRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateProductRequest) GetStatus() ProductStatus {
	if x != nil {
		return x.Status
	}
	return ProductStatus_ACTIVE
}

func (x *CreateProductRequest) GetAvailability() AvailabilityStatus {
	if x != nil {
		return x.Availability
	}
	return AvailabilityStatus_AVAILABLE
}

func (x *CreateProductRequest) GetStockQuantity() int32 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

type CreateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	mi := &file_product_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	DiscountPrice float64                `protobuf:"fixed64,5,opt,name=discount_price,json=discountPrice,proto3" json:"discount_price,omitempty"`
	Currency      string                 `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Status        ProductStatus          `protobuf:"varint,7,opt,name=status,proto3,enum=product.ProductStatus" json:"status,omitempty"`
	Availability  AvailabilityStatus     `protobuf:"varint,8,opt,name=availability,proto3,enum=product.AvailabilityStatus" json:"availability,omitempty"`
	StockQuantity int32                  `protobuf:"varint,9,opt,name=stock_quantity,json=stockQuantity,proto3" json:"stock_quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_product_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetDiscountPrice() float64 {
	if x != nil {
		return x.DiscountPrice
	}
	return 0
}

func (x *UpdateProductRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *UpdateProductRequest) GetStatus() ProductStatus {
	if x != nil {
		return x.Status
	}
	return ProductStatus_ACTIVE
}

func (x *UpdateProductRequest) GetAvailability() AvailabilityStatus {
	if x != nil {
		return x.Availability
	}
	return AvailabilityStatus_AVAILABLE
}

func (x *UpdateProductRequest) GetStockQuantity() int32 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

type UpdateProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductResponse) Reset() {
	*x = UpdateProductResponse{}
	mi := &file_product_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductResponse) ProtoMessage() {}

func (x *UpdateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetProductByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductByIdRequest) Reset() {
	*x = GetProductByIdRequest{}
	mi := &file_product_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRequest) ProtoMessage() {}

func (x *GetProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProductByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductByIdResponse) Reset() {
	*x = GetProductByIdResponse{}
	mi := &file_product_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdResponse) ProtoMessage() {}

func (x *GetProductByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdResponse.ProtoReflect.Descriptor instead.
func (*GetProductByIdResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductByIdResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_product_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductResponse) Reset() {
	*x = DeleteProductResponse{}
	mi := &file_product_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResponse) ProtoMessage() {}

func (x *DeleteProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          *int32                 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit         *int32                 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_product_product_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{9}
}

func (x *ListProductsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListProductsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_product_product_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{10}
}

func (x *ListProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ListProductsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListProductsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_product_product_proto protoreflect.FileDescriptor

const file_product_product_proto_rawDesc = "" +
	"\n" +
	"\x15product/product.proto\x12\aproduct\"\xfe\x02\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x01R\x05price\x12%\n" +
	"\x0ediscount_price\x18\x06 \x01(\x01R\rdiscountPrice\x12\x1a\n" +
	"\bcurrency\x18\a \x01(\tR\bcurrency\x12.\n" +
	"\x06status\x18\b \x01(\x0e2\x16.product.ProductStatusR\x06status\x12?\n" +
	"\favailability\x18\t \x01(\x0e2\x1b.product.AvailabilityStatusR\favailability\x12%\n" +
	"\x0estock_quantity\x18\n" +
	" \x01(\x05R\rstockQuantity\x12\x1d\n" +
	"\n" +
	"created_at\x18\v \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\f \x01(\tR\tupdatedAt\"\xbd\x02\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12%\n" +
	"\x0ediscount_price\x18\x05 \x01(\x01R\rdiscountPrice\x12\x1a\n" +
	"\bcurrency\x18\a \x01(\tR\bcurrency\x12.\n" +
	"\x06status\x18\b \x01(\x0e2\x16.product.ProductStatusR\x06status\x12?\n" +
	"\favailability\x18\t \x01(\x0e2\x1b.product.AvailabilityStatusR\favailability\x12%\n" +
	"\x0estock_quantity\x18\n" +
	" \x01(\x05R\rstockQuantity\"A\n" +
	"\x15CreateProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\xcd\x02\n" +
	"\x14UpdateProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12%\n" +
	"\x0ediscount_price\x18\x05 \x01(\x01R\rdiscountPrice\x12\x1a\n" +
	"\bcurrency\x18\x06 \x01(\tR\bcurrency\x12.\n" +
	"\x06status\x18\a \x01(\x0e2\x16.product.ProductStatusR\x06status\x12?\n" +
	"\favailability\x18\b \x01(\x0e2\x1b.product.AvailabilityStatusR\favailability\x12%\n" +
	"\x0estock_quantity\x18\t \x01(\x05R\rstockQuantity\"A\n" +
	"\x15UpdateProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"'\n" +
	"\x15GetProductByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"D\n" +
	"\x16GetProductByIdResponse\x12*\n" +
	"\aproduct\x18\x01 \x01(\v2\x10.product.ProductR\aproduct\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"1\n" +
	"\x15DeleteProductResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\\\n" +
	"\x13ListProductsRequest\x12\x17\n" +
	"\x04page\x18\x01 \x01(\x05H\x00R\x04page\x88\x01\x01\x12\x19\n" +
	"\x05limit\x18\x02 \x01(\x05H\x01R\x05limit\x88\x01\x01B\a\n" +
	"\x05_pageB\b\n" +
	"\x06_limit\"\x84\x01\n" +
	"\x14ListProductsResponse\x12,\n" +
	"\bproducts\x18\x01 \x03(\v2\x10.product.ProductR\bproducts\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\x12\x12\n" +
	"\x04page\x18\x03 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x04 \x01(\x05R\x05limit*M\n" +
	"\rProductStatus\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x00\x12\f\n" +
	"\bINACTIVE\x10\x01\x12\x10\n" +
	"\fOUT_OF_STOCK\x10\x02\x12\x10\n" +
	"\fDISCONTINUED\x10\x03*@\n" +
	"\x12AvailabilityStatus\x12\r\n" +
	"\tAVAILABLE\x10\x00\x12\r\n" +
	"\tBACKORDER\x10\x01\x12\f\n" +
	"\bPREORDER\x10\x022\xa0\x03\n" +
	"\x0eProductService\x12N\n" +
	"\rCreateProduct\x12\x1d.product.CreateProductRequest\x1a\x1e.product.CreateProductResponse\x12N\n" +
	"\rUpdateProduct\x12\x1d.product.UpdateProductRequest\x1a\x1e.product.UpdateProductResponse\x12Q\n" +
	"\x0eGetProductById\x12\x1e.product.GetProductByIdRequest\x1a\x1f.product.GetProductByIdResponse\x12N\n" +
	"\rDeleteProduct\x12\x1d.product.DeleteProductRequest\x1a\x1e.product.DeleteProductResponse\x12K\n" +
	"\fListProducts\x12\x1c.product.ListProductsRequest\x1a\x1d.product.ListProductsResponseB?Z=github.com/sahilrana7582/orderX/pkg/generated/product;productb\x06proto3"

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData []byte
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_product_proto_rawDesc), len(file_product_product_proto_rawDesc)))
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_product_product_proto_goTypes = []any{
	(ProductStatus)(0),             // 0: product.ProductStatus
	(AvailabilityStatus)(0),        // 1: product.AvailabilityStatus
	(*Product)(nil),                // 2: product.Product
	(*CreateProductRequest)(nil),   // 3: product.CreateProductRequest
	(*CreateProductResponse)(nil),  // 4: product.CreateProductResponse
	(*UpdateProductRequest)(nil),   // 5: product.UpdateProductRequest
	(*UpdateProductResponse)(nil),  // 6: product.UpdateProductResponse
	(*GetProductByIdRequest)(nil),  // 7: product.GetProductByIdRequest
	(*GetProductByIdResponse)(nil), // 8: product.GetProductByIdResponse
	(*DeleteProductRequest)(nil),   // 9: product.DeleteProductRequest
	(*DeleteProductResponse)(nil),  // 10: product.DeleteProductResponse
	(*ListProductsRequest)(nil),    // 11: product.ListProductsRequest
	(*ListProductsResponse)(nil),   // 12: product.ListProductsResponse
}
var file_product_product_proto_depIdxs = []int32{
	0,  // 0: product.Product.status:type_name -> product.ProductStatus
	1,  // 1: product.Product.availability:type_name -> product.AvailabilityStatus
	0,  // 2: product.CreateProductRequest.status:type_name -> product.ProductStatus
	1,  // 3: product.CreateProductRequest.availability:type_name -> product.AvailabilityStatus
	0,  // 4: product.UpdateProductRequest.status:type_name -> product.ProductStatus
	1,  // 5: product.UpdateProductRequest.availability:type_name -> product.AvailabilityStatus
	2,  // 6: product.GetProductByIdResponse.product:type_name -> product.Product
	2,  // 7: product.ListProductsResponse.products:type_name -> product.Product
	3,  // 8: product.ProductService.CreateProduct:input_type -> product.CreateProductRequest
	5,  // 9: product.ProductService.UpdateProduct:input_type -> product.UpdateProductRequest
	7,  // 10: product.ProductService.GetProductById:input_type -> product.GetProductByIdRequest
	9,  // 11: product.ProductService.DeleteProduct:input_type -> product.DeleteProductRequest
	11, // 12: product.ProductService.ListProducts:input_type -> product.ListProductsRequest
	4,  // 13: product.ProductService.CreateProduct:output_type -> product.CreateProductResponse
	6,  // 14: product.ProductService.UpdateProduct:output_type -> product.UpdateProductResponse
	8,  // 15: product.ProductService.GetProductById:output_type -> product.GetProductByIdResponse
	10, // 16: product.ProductService.DeleteProduct:output_type -> product.DeleteProductResponse
	12, // 17: product.ProductService.ListProducts:output_type -> product.ListProductsResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	file_product_product_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_product_proto_rawDesc), len(file_product_product_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		EnumInfos:         file_product_product_proto_enumTypes,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}
