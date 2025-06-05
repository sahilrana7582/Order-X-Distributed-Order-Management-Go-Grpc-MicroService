package domains

type OrderStatus string
type PaymentStatus string
type ShippingMethod string

const (
	// OrderStatus values
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusConfirmed  OrderStatus = "CONFIRMED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusShipped    OrderStatus = "SHIPPED"
	OrderStatusDelivered  OrderStatus = "DELIVERED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
	OrderStatusReturned   OrderStatus = "RETURNED"

	// PaymentStatus values
	PaymentStatusUnpaid   PaymentStatus = "UNPAID"
	PaymentStatusPaid     PaymentStatus = "PAID"
	PaymentStatusRefunded PaymentStatus = "REFUNDED"
	PaymentStatusFailed   PaymentStatus = "FAILED"

	// ShippingMethod values
	ShippingMethodStandard ShippingMethod = "STANDARD"
	ShippingMethodExpress  ShippingMethod = "EXPRESS"
	ShippingMethodSameDay  ShippingMethod = "SAME_DAY"
	ShippingMethodPickup   ShippingMethod = "PICKUP"
)

type Order struct {
	ID             string         `json:"id"`
	CustomerID     string         `json:"customer_id"`
	OrderNumber    string         `json:"order_number"`
	OrderStatus    OrderStatus    `json:"order_status"`
	PaymentStatus  PaymentStatus  `json:"payment_status"`
	ShippingMethod ShippingMethod `json:"shipping_method"`
	TotalAmount    float64        `json:"total_amount"`
	Currency       string         `json:"currency"`
}

type CreateOrderRequest struct {
	CustomerID     string         `json:"customer_id" validate:"required"`
	OrderNumber    string         `json:"order_number" validate:"required"`
	OrderStatus    OrderStatus    `json:"order_status"`
	PaymentStatus  PaymentStatus  `json:"payment_status"`
	ShippingMethod ShippingMethod `json:"shipping_method"`
	TotalAmount    float64        `json:"total_amount" validate:"required,gt=0"`
	Currency       string         `json:"currency"`
}

type CreateOrderResponse struct {
	ID      string `json:"id"`
	Message string `json:"message" validate:"omitempty,max=255"`
}

type GetAllOrdersResponse struct {
	Orders []Order `json:"orders"`
}
