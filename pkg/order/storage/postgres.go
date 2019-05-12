package storage

import (
	"fmt"

	"github.com/diskordanz/consumer/pkg/order/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	ordersCollection    = "orders"
	orderItemCollection = "order_items"
)

type UpperOrderStorage struct {
	db sqlbuilder.Database
}

func NewUpperOrderStorage(db sqlbuilder.Database) UpperOrderStorage {
	return UpperOrderStorage{db: db}
}

func (s UpperOrderStorage) ListItems(id int) ([]model.OrderItem, error) {
	orders := s.db.Collection(orderItemCollection)
	paginator := orders.Find(db.Cond{"order_id": id}).OrderBy("id")
	var result []model.OrderItem
	if err := paginator.All(&result); err != nil {
		return []model.OrderItem{}, err
	}
	return result, nil
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
	paginator := orders.Find().Where("consumer_id=?", id).OrderBy("id").Paginate(uint(count))
	var result []model.Order
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		fmt.Println(err)

		return []model.Order{}, err
	}
	return result, nil
}

func (s UpperOrderStorage) CreateOrderItem(item model.OrderItem) (model.OrderItem, error) {
	_, err := s.db.InsertInto(orderItemCollection).Columns("order_id", "product_id", "count").
		Values(item.OrderID, item.ProductID, item.Count).
		Exec()
	if err != nil {
		return model.OrderItem{}, err
	}
	return item, nil
}

func (s UpperOrderStorage) CreateOrder(item model.Order) (model.Order, error) {
	_, err := s.db.InsertInto(ordersCollection).Columns("consumer_id", "franchise_id", "time", "total", "status").
		Values(item.ConsumerID, item.FranchiseID, item.Time, item.Total, item.Status).
		Exec()
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	result := s.db.Collection(ordersCollection).Find().Where("time=? AND consumer_id=?", item.Time, item.ConsumerID)
	err = result.One(&order)
	return order, nil
}
