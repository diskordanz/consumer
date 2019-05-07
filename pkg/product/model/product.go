package model

type Product struct {
	ID          uint64
	Name        string
	Description string
	Image       string
	FranchiseID uint64
	Count       uint32
	Price       float32
	CategoryID  uint64
}
