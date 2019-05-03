package storage

import (
	"github.com/diskordanz/consumer/pkg/location/model"
)

type LocationStorage interface {
	// Get list of size *count* of Locations starting from *offset*
	ListOfLocations(id, count, offset int) ([]model.Location, error)
	ListOfLocationsByName(id, count, offset int, name string) ([]model.Location, error)

	Get(int) (model.Location, error)
}
