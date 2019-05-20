package storage

import (
	"github.com/diskordanz/web-consumer/pkg/franchise/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	franchisesCollection = "franchises"
)

type UpperFranchiseStorage struct {
	db sqlbuilder.Database
}

func NewUpperFranchiseStorage(db sqlbuilder.Database) UpperFranchiseStorage {
	return UpperFranchiseStorage{db: db}
}

func (s UpperFranchiseStorage) List(count, offset int) ([]model.Franchise, error) {
	franchises := s.db.Collection(franchisesCollection)
	paginator := franchises.Find().OrderBy("id").Paginate(uint(count))
	var result []model.Franchise
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.Franchise{}, err
	}
	return result, nil
}

func (s UpperFranchiseStorage) Get(id int) (model.Franchise, error) {
	result := s.db.Collection(franchisesCollection).Find(db.Cond{"id": id})
	var franchise model.Franchise
	err := result.One(&franchise)
	if err != nil {
		return model.Franchise{}, err
	}

	return franchise, nil
}
