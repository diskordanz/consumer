package storage

import (
	"github.com/diskordanz/consumer/pkg/product/model"
)

type ProductStorage interface {
	// Get list of size *count* of Products starting from *offset*
	List(count, offset int) ([]model.Product, error)
	Get(int) (model.Product, error)
}
