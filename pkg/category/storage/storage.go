package storage

import (
	"github.com/diskordanz/web-consumer/pkg/category/model"
)

type CategoryStorage interface {
	List() ([]model.Category, error)
}
