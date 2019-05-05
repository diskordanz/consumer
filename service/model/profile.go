package model

import (
	pkgProfileModel "github.com/diskordanz/consumer/pkg/profile/model"
)

type Consumer struct {
	ID          uint64        `json:"id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	PhoneNumber string        `json:"phone_number"`
	City        string        `json:"city"`
	Adress      string        `json:"adress"`
	Login       string        `json:"login"`
	Mail        string        `json:"mail"`
	Password    string        `json:"password"`
	Cart        map[int64]int `json:"cart"`
}

func MapToProfile(originProfile pkgProfileModel.Consumer) Consumer {
	return Consumer{
		ID:          originProfile.ID,
		FirstName:   originProfile.FirstName,
		LastName:    originProfile.LastName,
		PhoneNumber: originProfile.PhoneNumber,
		City:        originProfile.City,
		Adress:      originProfile.Adress,
		Login:       originProfile.Login,
		Mail:        originProfile.Mail,
		Password:    originProfile.Password,
		Cart:        originProfile.Cart,
	}
}

func MapToProfileList(originList []pkgProfileModel.Consumer) []Consumer {
	resultList := make([]Profile, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToProfile(v)
	}
	return resultList
}
