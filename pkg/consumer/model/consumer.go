package model

type Consumer struct {
	ID          uint64 `db:"id,omitempty"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	PhoneNumber string `db:"phone_number"`
	City        string `db:"city"`
	Address     string `db:"address"`
	Login       string `db:"login"`
	Mail        string `db:"mail"`
	Password    string `db:"password"`
}

type CartItem struct {
	ID         uint64 `db:"id"`
	ConsumerID uint64 `db:"consumer_id"`
	ProductID  uint64 `db:"product_id"`
	Count      uint   `db:"count"`
}
