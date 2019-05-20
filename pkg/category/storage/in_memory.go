package storage

import (
	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/web-consumer/pkg/category/model"
)

type InMemoryCategoryStorage struct {
	db []model.Category
}

func NewInMemoryCategoryStorage() InMemoryCategoryStorage {
	db := []model.Category{
		model.Category{ID: 1, Name: "Все категории"},
		model.Category{ID: 2, Name: "Продукты"},
		model.Category{ID: 3, Name: "Товары для детей"},
		model.Category{ID: 4, Name: "Косметика, бытовая химия"},
		model.Category{ID: 5, Name: "Техника, электроника"},
		model.Category{ID: 6, Name: "Товары для дома"},
		model.Category{ID: 7, Name: "Одежда, обувь, акскссуары"},
		model.Category{ID: 8, Name: "Зоотовары"},
		model.Category{ID: 9, Name: "Товары для авто"},
		model.Category{ID: 10, Name: "Спорт и отдых"},
		model.Category{ID: 11, Name: "Товары для дачи"},
		model.Category{ID: 12, Name: "Подарки, сувениры"},
	}
	return InMemoryCategoryStorage{db: db}
}

func (s InMemoryCategoryStorage) List() ([]model.Category, error) {
	result := linq.From(s.db)
	var slice []model.Category
	result.ToSlice(&slice)
	return slice, nil
}
