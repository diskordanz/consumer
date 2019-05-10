package storage

import (
	"github.com/diskordanz/consumer/pkg/consumer/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	consumersCollection = "consumers"
	cartCollection      = "cart_items"
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

func (s UpperConsumerStorage) Cart(id, count, offset int) ([]model.CartItem, error) {
	cart := s.db.Collection(cartCollection)
	paginator := cart.Find(db.Cond{"consumer_id": id}).OrderBy("id").Paginate(uint(count))
	var result []model.CartItem
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.CartItem{}, err
	}
	return result, nil
}

func (s UpperConsumerStorage) CreateCartItem(item model.CartItem) (model.CartItem, error) {
	_, err := s.db.InsertInto(cartCollection).Columns("consumer_id", "product_id", "count").
		Values(item.ConsumerID, item.ProductID, item.Count).
		Exec()
	if err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}

func (s UpperConsumerStorage) UpdateCartItem(item model.CartItem) (model.CartItem, error) {
	_, err := s.db.Update(cartCollection).
		Set("count", item.Count).
		Where("id=?", item.ID).
		Exec()
	if err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}

func (s UpperConsumerStorage) GetCartItem(id, productID int) (model.CartItem, error) {
	result := s.db.Collection(cartCollection).Find().Where("consumer_id=? AND product_id=?", id, productID)
	var item model.CartItem
	err := result.One(&item)
	if err != nil {
		return model.CartItem{}, err
	}
	return item, nil
}
