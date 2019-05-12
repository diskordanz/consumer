package model

import "time"

type Order struct {
	ID          uint64    `db:"id,omitempty"`
	ConsumerID  uint64    `db:"consumer_id"`
	FranchiseID uint64    `db:"franchise_id"`
	Time        time.Time `db:"time"`
	Total       float32   `db:"total"`
	Status      string    `db:"status"`
}

type OrderItem struct {
	ID        uint64 `db:"id,omitempty"`
	OrderID   uint64 `db:"order_id"`
	ProductID uint64 `db:"product_id"`
	Count     uint   `db:"count"`
}
