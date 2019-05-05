package model

type Product struct {
	ID          uint64
	Name        string
	Description string
	Image       string
	FranchiseID uint64
	Count       int32
	Price       float32
}
