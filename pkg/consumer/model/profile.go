package model

type Consumer struct {
	ID          uint64
	FirstName   string
	LastName    string
	PhoneNumber string
	City        string
	Adress      string
	Login       string
	Mail        string
	Password    string
	Cart        map[int64]int
}
