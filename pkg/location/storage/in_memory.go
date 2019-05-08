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
		model.Location{ID: 1, Type: "магазин", City: "Гомель", FranchiseID: 1, Locality: "Березки", Adress: "ул. 60 лет СССР, 10В", OpeningHours: "08:00-23:00"},
		model.Location{ID: 2, Type: "магазин", City: "Гомель", FranchiseID: 1, Locality: "Большевик", Adress: "ул.Советская, 36", OpeningHours: "08:00-22:00"},
		model.Location{ID: 3, Type: "магазин", City: "Гомель", FranchiseID: 1, Locality: "Брагин", Adress: "ул.Советская, 21", OpeningHours: "08:00-22:00"},
		model.Location{ID: 4, Type: "магазин", City: "Гомель", FranchiseID: 1, Locality: "Гомель", Adress: "пр-т Речицкий, 5в", OpeningHours: "09:00-23:00"},
		model.Location{ID: 5, Type: "магазин", City: "Гомель", FranchiseID: 2, Locality: "Гомель", Adress: "ул. Косарева 18", OpeningHours: "09:00-23:00"},
		model.Location{ID: 6, Type: "магазин", City: "Минск", FranchiseID: 2, Locality: "Минск", Adress: "Игуменский тракт, 30", OpeningHours: "09:00-23:00"},
	}
	return InMemoryLocationStorage{db: db}
}

func (s InMemoryLocationStorage) List(id, count, offset int) ([]model.Location, error) {
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

func (s InMemoryLocationStorage) ListOfLocationsByName(id, count, offset int, name string) ([]model.Location, error) {
	result := linq.From(s.db).Where(func(c interface{}) bool {
		return strings.Contains(strings.ToLower((c.(model.Location).City)), strings.ToLower(name)) && c.(model.Location).FranchiseID == id
	}).Skip(offset).Take(count)
	var slice []model.Location
	result.ToSlice(&slice)
	return slice, nil
}
