package model

import (
	pkgProfileModel "github.com/diskordanz/consumer/pkg/profile/model"
)

type Profile struct {
	ID        uint64 `json:"id"`
	CountryID int32  `json:"country_id"`
	Name      string `json:"name"`
}

func MapToProfile(originProfile pkgProfileModel.Profile) Profile {
	return Profile{ID: originProfile.ID, Name: originProfile.Name}
}

func MapToProfileList(originList []pkgProfileModel.Profile) []Profile {
	resultList := make([]Profile, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToProfile(v)
	}
	return resultList
}
