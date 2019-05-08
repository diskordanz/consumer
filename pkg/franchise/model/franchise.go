package model

type Franchise struct {
	ID          uint64 `db:"id"`
	CountryID   int    `db:"country_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Image       string `db:"image"`
}
