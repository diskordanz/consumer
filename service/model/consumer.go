package model

import (
	pkgConsumerModel "github.com/diskordanz/consumer/pkg/consumer/model"
)

type Consumer struct {
	ID          uint64        `json:"id"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	PhoneNumber string        `json:"phone_number"`
	City        string        `json:"city"`
	Address     string        `json:"address"`
	Login       string        `json:"login"`
	Mail        string        `json:"mail"`
	Password    string        `json:"password"`
	Cart        map[int64]int `json:"cart"`
}

func MapToConsumer(originConsumer pkgConsumerModel.Consumer) Consumer {
	return Consumer{
		ID:          originConsumer.ID,
		FirstName:   originConsumer.FirstName,
		LastName:    originConsumer.LastName,
		PhoneNumber: originConsumer.PhoneNumber,
		City:        originConsumer.City,
		Address:     originConsumer.Address,
		Login:       originConsumer.Login,
		Mail:        originConsumer.Mail,
		Password:    originConsumer.Password,
		Cart:        originConsumer.Cart,
	}
}

func MapToConsumerDB(originConsumer Consumer) pkgConsumerModel.Consumer {
	return pkgConsumerModel.Consumer{
		ID:          originConsumer.ID,
		FirstName:   originConsumer.FirstName,
		LastName:    originConsumer.LastName,
		PhoneNumber: originConsumer.PhoneNumber,
		City:        originConsumer.City,
		Address:     originConsumer.Address,
		Login:       originConsumer.Login,
		Mail:        originConsumer.Mail,
		Password:    originConsumer.Password,
		Cart:        originConsumer.Cart,
	}
}

func MapToConsumerList(originList []pkgConsumerModel.Consumer) []Consumer {
	resultList := make([]Consumer, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToConsumer(v)
	}
	return resultList
}
