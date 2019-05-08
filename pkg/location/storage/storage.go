package storage

import (
	"github.com/diskordanz/consumer/pkg/location/model"
)

type LocationStorage interface {
	//ListOfLocationsByName(id, count, offset int, name string) ([]model.Location, error)
	List(id, count, offset int) ([]model.Location, error)
	Get(int) (model.Location, error)
}
