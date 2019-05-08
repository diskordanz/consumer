package model

type Consumer struct {
	ID          uint64        `db:"id"`
	FirstName   string        `db:"first_name"`
	LastName    string        `db:"last_name"`
	PhoneNumber string        `db:"phone_number"`
	City        string        `db:"city"`
	Address     string        `db:"address"`
	Login       string        `db:"login"`
	Mail        string        `db:"mail"`
	Password    string        `db:"password"`
	Cart        map[int64]int `db:"cart"`
}
