package storage

import (
	"github.com/diskordanz/web-consumer/pkg/product/model"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	productsCollection = "products"
)

type UpperProductStorage struct {
	db sqlbuilder.Database
}

func NewUpperProductStorage(db sqlbuilder.Database) UpperProductStorage {
	return UpperProductStorage{db: db}
}

func (s UpperProductStorage) List(name string, count, offset int) ([]model.Product, error) {
	products := s.db.Collection(productsCollection)
	paginator := products.Find().Where("name ILIKE ?", "%%"+name+"%%").OrderBy("id").Paginate(uint(count))
	var result []model.Product
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.Product{}, err
	}
	return result, nil
}

func (s UpperProductStorage) Get(id int) (model.Product, error) {
	result := s.db.Collection(productsCollection).Find(db.Cond{"id": id})
	var product model.Product
	err := result.One(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (s UpperProductStorage) ListByCategory(id uint64, name string, count, offset int) ([]model.Product, error) {
	products := s.db.Collection(productsCollection)
	paginator := products.Find().Where("name ILIKE ? AND category_id=?", "%%"+name+"%%", id).OrderBy("id").Paginate(uint(count))
	var result []model.Product
	if err := paginator.Page(uint(offset/count + 1)).All(&result); err != nil {
		return []model.Product{}, err
	}
	return result, nil
}
