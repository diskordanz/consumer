package storage

import (
	"errors"

	"github.com/diskordanz/consumer/pkg/profile/model"
)

type InMemoryProfileStorage struct {
	db []model.Consumer
}

func NewInMemoryProfileStorage() InMemoryProfileStorage {
	db := []model.Consumer{
		model.Consumer{ID: 1, FirstName: "Надежда", LastName: "Барсукова", PhoneNumber: "+375(99)99-99-999", City: "Гомель", Adress: "Ерёмино", Login: "nadya", Mail: "nadya@gmail.com", Password: "1111"},
		model.Consumer{ID: 2, FirstName: "Consumer2", LastName: "", PhoneNumber: "", City: "", Adress: "", Login: "consumer2", Mail: "", Password: "2222"},
	}
	return InMemoryProfileStorage{db: db}
}

func (s InMemoryProfileStorage) Get(id int) (model.Consumer, error) {
	if id > len(s.db) || id < 0 {
		return model.Consumer{}, errors.New("error. no such entry in database")
	}
	return s.db[id], nil
}
