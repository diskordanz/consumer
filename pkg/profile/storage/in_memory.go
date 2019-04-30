package storage

import (
	"errors"

	"github.com/diskordanz/consumer/pkg/profile/model"
)

type InMemoryProfileStorage struct {
	db []model.Profile
}

func NewInMemoryProfileStorage() InMemoryProfileStorage {
	db := []model.Profile{
		model.Profile{ID: 1, Name: "Evroopt"},
		model.Profile{ID: 2, Name: "Kopeechka"},
		model.Profile{ID: 3, Name: "Dobryk"},
	}
	return InMemoryProfileStorage{db: db}
}

func (s InMemoryProfileStorage) Get(id int) (model.Profile, error) {
	if id > len(s.db) || id < 0 {
		return model.Profile{}, errors.New("error. no such entry in database")
	}
	return s.db[id], nil
}
