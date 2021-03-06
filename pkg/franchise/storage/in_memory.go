package storage

import (
	"errors"
	"strings"

	"github.com/ahmetb/go-linq"
	"github.com/diskordanz/web-consumer/pkg/franchise/model"
)

type InMemoryFranchiseStorage struct {
	db []model.Franchise
}

func NewInMemoryFranchiseStorage() InMemoryFranchiseStorage {
	db := []model.Franchise{
		model.Franchise{ID: 1, CountryID: 1, Name: "Евроопт", Description: "Евроопт - сеть магазинов по всей Беларуси. Продукты питания и потребительские товары.", Image: "logo_evroopt.jpg"},
		model.Franchise{ID: 2, CountryID: 1, Name: "Гиппо", Description: "Гиппо - сеть гипермаркетов, супермаркетов. Продукты питания, товары для дома, изделия собственного производства.", Image: "logo_gippo.jpg"},
		model.Franchise{ID: 3, CountryID: 2, Name: "Доброном", Description: "", Image: "logo_dobronom.jpg"},
		model.Franchise{ID: 4, CountryID: 2, Name: "5 Элемент", Description: "5 элемент - продажа бытовой техники и аудио-видео в Беларуси. Официальное дилерство ведущих мировых производителей.", Image: "logo_5element.png"},
		model.Franchise{ID: 5, CountryID: 2, Name: "Электросила", Description: "Электросила - оптовая и розничная торговля бытовой техникой и электроникой.", Image: "logo_electrosila.jpg"},
		model.Franchise{ID: 6, CountryID: 2, Name: "Карусель", Description: "", Image: "logo_karusel.png"},
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
