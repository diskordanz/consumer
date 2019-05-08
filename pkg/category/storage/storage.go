package storage

import (
	"github.com/diskordanz/consumer/pkg/category/model"
)

type CategoryStorage interface {
	List() ([]model.Category, error)
}
