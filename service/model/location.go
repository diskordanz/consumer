package model

import (
	pkgLocationModel "github.com/diskordanz/consumer/pkg/location/model"
)

type Location struct {
	ID          uint64 `json:"id"`
	FranchiseID int    `json:"franchise_id"`
	Name        string `json:"name"`
}

func MapToLocation(originLocation pkgLocationModel.Location) Location {
	return Location{ID: originLocation.ID, Name: originLocation.Name}
}

func MapToLocationsList(originList []pkgLocationModel.Location) []Location {
	resultList := make([]Location, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToLocation(v)
	}
	return resultList
}
