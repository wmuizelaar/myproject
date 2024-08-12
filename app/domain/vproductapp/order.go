package vproductapp

import (
	"github.com/wmuizelaar/myproject/business/domain/vproductbus"
	"github.com/wmuizelaar/myproject/business/sdk/order"
)

var defaultOrderBy = order.NewBy("product_id", order.ASC)

var orderByFields = map[string]string{
	"product_id": vproductbus.OrderByProductID,
	"user_id":    vproductbus.OrderByUserID,
	"name":       vproductbus.OrderByName,
	"cost":       vproductbus.OrderByCost,
	"quantity":   vproductbus.OrderByQuantity,
	"user_name":  vproductbus.OrderByUserName,
}
