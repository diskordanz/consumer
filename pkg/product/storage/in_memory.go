package storage

import (
	"errors"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/consumer/pkg/product/model"
)

type InMemoryProductStorage struct {
	db []model.Product
}

func NewInMemoryProductStorage() InMemoryProductStorage {
	db := []model.Product{
		model.Product{ID: 1, CategoryID: 2, Name: "Сок Добрый", Description: "Очень добрый", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 2, CategoryID: 2, Name: "Продукт", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 3, CategoryID: 5, Name: "Техника", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 4, CategoryID: 6, Name: "6", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
	}
	return InMemoryProductStorage{db: db}
}

func (s InMemoryProductStorage) List(name string, count, offset int) ([]model.Product, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Product).Name)), strings.ToLower(name))
	}).Skip(offset).Take(count)

	var slice []model.Product
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryProductStorage) ListByCategory(id uint64, name string, count, offset int) ([]model.Product, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Product).Name)), strings.ToLower(name))
	}).Where(func(c interface{}) bool {
		return c.(model.Product).CategoryID == id
	}).Skip(offset).Take(count)

	var slice []model.Product
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryProductStorage) Get(id int) (model.Product, error) {
	if id > len(s.db) || id < 0 {
		return model.Product{}, errors.New("error. no such entry in database")
	}
	return s.db[id-1], nil
}
