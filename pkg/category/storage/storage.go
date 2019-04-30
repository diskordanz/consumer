package storage

import (
	"github.com/diskordanz/consumer/pkg/category/model"
)

type CategoryStorage interface {
	// Get list of size *count* of Categorys starting from *offset*
	List() ([]model.Category, error)
}
