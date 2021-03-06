package model

type Location struct {
	ID           uint64 `db:"id"`
	FranchiseID  int    `db:"franchise_id"`
	City         string `db:"city"`
	Locality     string `db:"locality"`
	Adress       string `db:"address"`
	OpeningHours string `db:"opening_hours"`
	Latitude     string `db:"latitude"`
	Longitude    string `db:"longitude"`
}
