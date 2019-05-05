package model

type Order struct {
	ID       uint64
	UserID   uint64
	Products map[int64]int
	Date     string
	Total    float32
	Status   string
}
