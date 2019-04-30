package storage

import (
	"github.com/diskordanz/consumer/pkg/order/model"
)

type OrderStorage interface {
	// Get list of size *count* of Orders starting from *offset*
	List(id, count, offset int) ([]model.Order, error)
	Get(consumerID, orderID int) (model.Order, error)
}
