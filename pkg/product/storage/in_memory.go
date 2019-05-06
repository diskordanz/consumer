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
		model.Product{ID: 1, Name: "Сок Добрый", Description: "Очень добрый", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 2, Name: "", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 3, Name: "", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
		model.Product{ID: 4, Name: "", Description: "", Image: "", FranchiseID: 1, Count: 0, Price: 0},
	}
	return InMemoryProductStorage{db: db}
}

func (s InMemoryProductStorage) List(name string, count, offset int) ([]model.Product, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Product).Name)), strings.ToLower(name))
	}).Skip(offset).Take(count)

	//result := linq.From(s.db).Skip(offset).Take(count)

	var slice []model.Product
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryProductStorage) Get(id int) (model.Product, error) {
	if id > len(s.db) || id < 0 {
		return model.Product{}, errors.New("error. no such entry in database")
	}
	return s.db[id], nil
}
