package storage

import (
	"errors"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/consumer/pkg/location/model"
)

type InMemoryLocationStorage struct {
	db []model.Location
}

func NewInMemoryLocationStorage() InMemoryLocationStorage {
	db := []model.Location{
		model.Location{ID: 1, Name: "1", Description: "Eto evrik", FranchiseID: 1},
		model.Location{ID: 2, Name: "2", Description: "Eto kopeechka", FranchiseID: 2},
		model.Location{ID: 3, Name: "3", Description: "Eto dobryk", FranchiseID: 2},
	}
	return InMemoryLocationStorage{db: db}
}

func (s InMemoryLocationStorage) ListOfFranchise(id, count, offset int) ([]model.Location, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return c.(model.Location).FranchiseID == id
	}).Skip(offset).Take(count)
	var slice []model.Location
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryLocationStorage) Get(id int) (model.Location, error) {
	if id > len(s.db) || id < 0 {
		return model.Location{}, errors.New("error. no such entry in database")
	}
	return s.db[id], nil
}

func (s InMemoryLocationStorage) ListOfFranchiseByName(id, count, offset int, name string) ([]model.Location, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Location).Name)), strings.ToLower(name)) && c.(model.Location).FranchiseID == id
	}).Skip(offset).Take(count)
	var slice []model.Location
	result.ToSlice(&slice)
	return slice, nil
}
