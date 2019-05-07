package storage

import (
	"github.com/diskordanz/consumer/pkg/consumer/model"
)

type ConsumerStorage interface {
	// Get list of size *count* of Consumers starting from *offset*
	Get(int) (model.Consumer, error)
	Create(model.Consumer) (model.Consumer, error)
	Update(model.Consumer) (model.Consumer, error)
}
