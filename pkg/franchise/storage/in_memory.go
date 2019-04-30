package storage

import (
	"errors"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/consumer/pkg/franchise/model"
)

type InMemoryFranchiseStorage struct {
	db []model.Franchise
}

func NewInMemoryFranchiseStorage() InMemoryFranchiseStorage {
	db := []model.Franchise{
		model.Franchise{ID: 1, CountryID: 1, Name: "Evroopt", Description: "Eto evrik"},
		model.Franchise{ID: 2, CountryID: 1, Name: "Kopeechka", Description: "Eto kopeechka"},
		model.Franchise{ID: 3, CountryID: 2, Name: "Dobryk", Description: "Eto dobryk"},
	}
	return InMemoryFranchiseStorage{db: db}
}

func (s InMemoryFranchiseStorage) List(count, offset int) ([]model.Franchise, error) {
	result := linq.From(s.db).Skip(offset).Take(count)
	var slice []model.Franchise
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryFranchiseStorage) Get(id int) (model.Franchise, error) {
	if id > len(s.db) || id < 0 {
		return model.Franchise{}, errors.New("error. no such entry in database")
	}
	return s.db[id-1], nil
}

func (s InMemoryFranchiseStorage) SearchFranchisesByName(count, offset int, name string) ([]model.Franchise, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Franchise).Name)), strings.ToLower(name))
	}).Skip(offset).Take(count)
	var slice []model.Franchise
	result.ToSlice(&slice)
	return slice, nil
}

func (s InMemoryFranchiseStorage) GetListByCountryId(countryID, count, offset int) ([]model.Franchise, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return c.(model.Franchise).CountryID == countryID
	}).Skip(offset).Take(count)
	var slice []model.Franchise
	result.ToSlice(&slice)
	return slice, nil
}
