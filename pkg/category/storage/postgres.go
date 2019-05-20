package storage

import (
	"github.com/diskordanz/web-consumer/pkg/category/model"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	categoriesCollection = "categories"
)

type UpperCategoryStorage struct {
	db sqlbuilder.Database
}

func NewUpperCategoryStorage(db sqlbuilder.Database) UpperCategoryStorage {
	return UpperCategoryStorage{db: db}
}

func (s UpperCategoryStorage) List() ([]model.Category, error) {
	categories := s.db.Collection(categoriesCollection)
	var result []model.Category
	if err := categories.Find().All(&result); err != nil {
		return []model.Category{}, err
	}
	return result, nil
}
