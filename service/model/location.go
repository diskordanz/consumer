package model

import (
	pkgLocationModel "github.com/diskordanz/web-consumer/pkg/location/model"
)

type Location struct {
	ID           uint64 `json:"id"`
	FranchiseID  int    `json:"franchise_id"`
	City         string `json:"city"`
	Locality     string `json:"locality"`
	Adress       string `json:"address"`
	OpeningHours string `json:"opening_hours"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

func MapToLocation(originLocation pkgLocationModel.Location) Location {
	return Location{
		ID:           originLocation.ID,
		City:         originLocation.City,
		FranchiseID:  originLocation.FranchiseID,
		Locality:     originLocation.Locality,
		Adress:       originLocation.Adress,
		OpeningHours: originLocation.OpeningHours,
		Latitude:     originLocation.Latitude,
		Longitude:    originLocation.Longitude,
	}
}

func MapToLocationsList(originList []pkgLocationModel.Location) []Location {
	resultList := make([]Location, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToLocation(v)
	}
	return resultList
}
