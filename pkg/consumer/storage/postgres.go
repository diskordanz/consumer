package storage

import (
	"github.com/diskordanz/consumer/pkg/consumer/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	consumersCollection = "consumers"
)

type UpperConsumerStorage struct {
	db sqlbuilder.Database
}

func NewUpperConsumerStorage(db sqlbuilder.Database) UpperConsumerStorage {
	return UpperConsumerStorage{db: db}
}

func (s UpperConsumerStorage) Create(consumer model.Consumer) (model.Consumer, error) {
	_, err := s.db.InsertInto(consumersCollection).Columns("first_name", "last_name", "address", "city", "login", "mail", "password", "phone_number").
		Values(consumer.FirstName, consumer.LastName, consumer.Address, consumer.City, consumer.Login, consumer.Mail, consumer.Password, consumer.PhoneNumber).
		Exec()
	if err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}

func (s UpperConsumerStorage) Update(consumer model.Consumer) (model.Consumer, error) {
	_, err := s.db.Update(consumersCollection).
		Set("first_name", consumer.FirstName, "last_name", consumer.LastName, "address", consumer.Address, "city", consumer.City, "login", consumer.Login, "mail", consumer.Mail, "password", consumer.Password, "phone_number", consumer.PhoneNumber).
		Where("id = ?", consumer.ID).
		Exec()

	if err != nil {
		return model.Consumer{}, err
	}
	return consumer, nil
}

func (s UpperConsumerStorage) Get(id int) (model.Consumer, error) {
	result := s.db.Collection(consumersCollection).Find(db.Cond{"id": id})
	var consumer model.Consumer
	err := result.One(&consumer)
	if err != nil {
		return model.Consumer{}, err
	}

	return consumer, nil
}
