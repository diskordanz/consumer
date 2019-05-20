package storage

import (
	"github.com/diskordanz/web-consumer/pkg/franchise/model"
)

type FranchiseStorage interface {
	List(count, offset int) ([]model.Franchise, error)
	Get(int) (model.Franchise, error)
	//SearchFranchisesByName(count, offset int, name string) ([]model.Franchise, error)
	//GetListByCountryId(countryID, count, offset int) ([]model.Franchise, error)
}
