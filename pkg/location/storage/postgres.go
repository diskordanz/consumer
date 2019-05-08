package storage

import (
	"github.com/diskordanz/consumer/pkg/location/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	locationsCollection = "locations"
)

type UpperLocationStorage struct {
	db sqlbuilder.Database
}

func NewUpperLocationStorage(db sqlbuilder.Database) UpperLocationStorage {
	return UpperLocationStorage{db: db}
}

func (s UpperLocationStorage) List(id, count, offset int) ([]model.Location, error) {
	locations := s.db.Collection(locationsCollection)
	paginator := locations.Find(db.Cond{"franchise_id": id}).OrderBy("id").Paginate(uint(count))
	var result []model.Location
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.Location{}, err
	}
	return result, nil
}

func (s UpperLocationStorage) Get(id int) (model.Location, error) {
	result := s.db.Collection(locationsCollection).Find(db.Cond{"id": id})
	var location model.Location
	err := result.One(&location)
	if err != nil {
		return model.Location{}, err
	}

	return location, nil
}
