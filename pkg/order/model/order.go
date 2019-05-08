package model

type Order struct {
	ID         uint64        `db:"id"`
	ConsumerID uint64        `db:"consumer_id"`
	Products   map[int64]int `db:"products"`
	Date       string        `db:"date"`
	Total      float32       `db:"total"`
	Status     string        `db:"status"`
}
