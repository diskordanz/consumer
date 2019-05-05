package model

import (
	pkgLocationModel "github.com/diskordanz/consumer/pkg/location/model"
)

type Location struct {
	ID           uint64 `json:"id"`
	FranchiseID  int    `json:"franchise_id"`
	City         string `json:"city"`
	Locality     string `json:"inhabited_locality"`
	Adress       string `json:"adress"`
	OpeningHours string `json:"opening_hours"`
	Type         string `json:"type"`
}

func MapToLocation(originLocation pkgLocationModel.Location) Location {
	return Location{
		ID:           originLocation.ID,
		City:         originLocation.City,
		FranchiseID:  originLocation.FranchiseID,
		Locality:     originLocation.Locality,
		Adress:       originLocation.Adress,
		OpeningHours: originLocation.OpeningHours,
	}
}

func MapToLocationsList(originList []pkgLocationModel.Location) []Location {
	resultList := make([]Location, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToLocation(v)
	}
	return resultList
}
