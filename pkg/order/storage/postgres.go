package storage

import (
	"github.com/diskordanz/consumer/pkg/order/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	ordersCollection = "orders"
)

type UpperOrderStorage struct {
	db sqlbuilder.Database
}

func NewUpperOrderStorage(db sqlbuilder.Database) UpperOrderStorage {
	return UpperOrderStorage{db: db}
}

func (s UpperOrderStorage) Get(id int) (model.Order, error) {
	result := s.db.Collection(ordersCollection).Find(db.Cond{"id": id})
	var order model.Order
	err := result.One(&order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (s UpperOrderStorage) List(id, count, offset int) ([]model.Order, error) {
	orders := s.db.Collection(ordersCollection)
	paginator := orders.Find(db.Cond{"consumer_id": id}).OrderBy("id").Paginate(uint(count))
	var result []model.Order
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.Order{}, err
	}
	return result, nil
}
