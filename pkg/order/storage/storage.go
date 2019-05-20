package storage

import (
	"github.com/diskordanz/web-consumer/pkg/order/model"
)

type OrderStorage interface {
	// Get list of size *count* of Orders starting from *offset*
	List(id, count, offset int) ([]model.Order, error)
	ListItems(id int) ([]model.OrderItem, error)
	Get(id int) (model.Order, error)

	CreateOrder(order model.Order) (model.Order, error)
	CreateOrderItem(item model.OrderItem) (model.OrderItem, error)
}
