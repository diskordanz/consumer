package storage

import (
	"errors"

	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/consumer/pkg/order/model"
)

type InMemoryOrderStorage struct {
	db []model.Order
}

func NewInMemoryOrderStorage() InMemoryOrderStorage {
	db := []model.Order{
		model.Order{ID: 1, ConsumerID: 1, Total: 0, Status: ""},
	}
	return InMemoryOrderStorage{db: db}
}

func (s InMemoryOrderStorage) List(id, count, offset int) ([]model.Order, error) {
	result := linq.From(s.db).Skip(offset).Take(count)
	var slice []model.Order
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryOrderStorage) Get(id int) (model.Order, error) {
	if id > len(s.db) || id < 0 {
		return model.Order{}, errors.New("error. no such entry in database")
	}
	return s.db[id], nil
}
