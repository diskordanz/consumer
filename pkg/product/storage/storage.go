package storage

import (
	"github.com/diskordanz/web-consumer/pkg/product/model"
)

type ProductStorage interface {
	// Get list of size *count* of Products starting from *offset*
	List(name string, count, offset int) ([]model.Product, error)
	ListByCategory(id uint64, name string, count, offset int) ([]model.Product, error)
	Get(int) (model.Product, error)
}
