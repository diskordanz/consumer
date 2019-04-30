package storage

import (
	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/consumer/pkg/category/model"
)

type InMemoryCategoryStorage struct {
	db []model.Category
}

func NewInMemoryCategoryStorage() InMemoryCategoryStorage {
	db := []model.Category{
		model.Category{ID: 1, Name: "food"},
		model.Category{ID: 2, Name: "not food"},
	}
	return InMemoryCategoryStorage{db: db}
}

func (s InMemoryCategoryStorage) List() ([]model.Category, error) {
	result := linq.From(s.db)
	var slice []model.Category
	result.ToSlice(&slice)
	return slice, nil
}
