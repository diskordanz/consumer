package model

type Product struct {
	ID          uint64  `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Image       string  `db:"image"`
	FranchiseID uint64  `db:"franchise_id"`
	Count       uint32  `db:"count"`
	Price       float32 `db:"price"`
	CategoryID  uint64  `db:"category_id"`
}
