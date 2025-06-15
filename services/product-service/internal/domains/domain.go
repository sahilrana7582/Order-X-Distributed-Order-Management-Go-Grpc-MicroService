package domains

import "time"

const (
	ProductStatusActive       = "ACTIVE"
	ProductStatusInactive     = "INACTIVE"
	ProductStatusOutOfStock   = "OUT_OF_STOCK"
	ProductStatusDiscontinued = "DISCONTINUED"

	AvailabilityInStock     = "IN_STOCK"
	AvailabilityPreOrder    = "PRE_ORDER"
	AvailabilityUnavailable = "UNAVAILABLE"
)

type Product struct {
	ID            string
	Name          string
	Description   string
	Price         float64
	DiscountPrice float64
	Currency      string
	Status        string
	Availability  string
	StockQuantity int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
