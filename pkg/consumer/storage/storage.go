package storage

import (
	"github.com/diskordanz/consumer/pkg/consumer/model"
)

type ConsumerStorage interface {
	Get(int) (model.Consumer, error)
	Create(model.Consumer) (model.Consumer, error)
	Update(model.Consumer) (model.Consumer, error)
	Cart(id, count, offset int) ([]model.CartItem, error)
}
