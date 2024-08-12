package vproductbus

import (
	"time"

	"github.com/wmuizelaar/myproject/business/domain/productbus"
	"github.com/wmuizelaar/myproject/business/domain/userbus"
	"github.com/google/uuid"
)

// Product represents an individual product with extended information.
type Product struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        productbus.Name
	Cost        float64
	Quantity    int
	DateCreated time.Time
	DateUpdated time.Time
	UserName    userbus.Name
}
