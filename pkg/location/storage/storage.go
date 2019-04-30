package storage

import (
	"github.com/diskordanz/consumer/pkg/location/model"
)

type LocationStorage interface {
	// Get list of size *count* of Locations starting from *offset*
	ListOfFranchise(id, count, offset int) ([]model.Location, error)
	ListOfFranchiseByName(id, count, offset int, name string) ([]model.Location, error)

	Get(int) (model.Location, error)
}
