package vproductbus

import (
	"github.com/wmuizelaar/myproject/business/domain/productbus"
	"github.com/wmuizelaar/myproject/business/domain/userbus"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID       *uuid.UUID
	Name     *productbus.Name
	Cost     *float64
	Quantity *int
	UserName *userbus.Name
}
