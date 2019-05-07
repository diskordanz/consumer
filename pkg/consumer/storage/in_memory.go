package storage

import (
	"errors"

	"github.com/diskordanz/consumer/pkg/consumer/model"
)

type InMemoryConsumerStorage struct {
	db []model.Consumer
}

func NewInMemoryConsumerStorage() InMemoryConsumerStorage {
	db := []model.Consumer{
		model.Consumer{ID: 1, FirstName: "Надежда", LastName: "Барсукова", PhoneNumber: "+375(99)99-99-999", City: "Гомель", Adress: "Ерёмино", Login: "nadya", Mail: "nadya@gmail.com", Password: "1111"},
		model.Consumer{ID: 2, FirstName: "Consumer2", LastName: "", PhoneNumber: "", City: "", Adress: "", Login: "consumer2", Mail: "", Password: "2222"},
	}
	return InMemoryConsumerStorage{db: db}
}

func (s InMemoryConsumerStorage) Get(id int) (model.Consumer, error) {
	if id > len(s.db) || id < 0 {
		return model.Consumer{}, errors.New("error. no such entry in database")
	}
	return s.db[id-1], nil
}

func (s InMemoryConsumerStorage) Create(consumer model.Consumer) (model.Consumer, error) {
	consumer.ID = uint64(len(s.db) + 1)
	s.db = append(s.db, consumer)
	return consumer, nil
}

func (s InMemoryConsumerStorage) Update(consumer model.Consumer) (model.Consumer, error) {
	s.db[consumer.ID-1] = consumer
	return consumer, nil
}
