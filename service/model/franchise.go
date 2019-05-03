package model

import (
	pkgFranchiseModel "github.com/diskordanz/consumer/pkg/franchise/model"
)

type Franchise struct {
	ID          uint64 `json:"id"`
	CountryID   int    `json:"country_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func MapToFranchise(originFranchise pkgFranchiseModel.Franchise) Franchise {
	return Franchise{ID: originFranchise.ID, Name: originFranchise.Name, CountryID: originFranchise.CountryID, Image: originFranchise.Image, Description: originFranchise.Description}
}

func MapToFranchiseList(originList []pkgFranchiseModel.Franchise) []Franchise {
	resultList := make([]Franchise, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToFranchise(v)
	}
	return resultList
}
