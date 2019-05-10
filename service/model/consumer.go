package model

import (
	pkgConsumerModel "github.com/diskordanz/consumer/pkg/consumer/model"
	pkgProductModel "github.com/diskordanz/consumer/pkg/product/model"
)

type Consumer struct {
	ID          uint64 `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Login       string `json:"login"`
	Mail        string `json:"mail"`
	Password    string `json:"password"`
}

type CartItem struct {
	ID         uint64  `json:"id"`
	ConsumerID uint64  `json:"consumer_id"`
	Product    Product `json:"product"`
	Count      uint    `json:"count"`
}

func MapToCart(originList []pkgConsumerModel.CartItem, originProducts []pkgProductModel.Product) []CartItem {
	resultList := make([]CartItem, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToCartItem(v, originProducts[i])
	}
	return resultList
}

func MapToCartItem(originItemCart pkgConsumerModel.CartItem, originProduct pkgProductModel.Product) CartItem {
	return CartItem{
		ID:         originItemCart.ID,
		ConsumerID: originItemCart.ConsumerID,
		Product: Product{
			ID:    originItemCart.ProductID,
			Name:  originProduct.Name,
			Image: originProduct.Image,
			Price: originProduct.Price,
		},
		Count: originItemCart.Count,
	}
}

func MapToCartItemDB(originItemCart CartItem) pkgConsumerModel.CartItem {
	return pkgConsumerModel.CartItem{
		ID:         originItemCart.ID,
		ConsumerID: originItemCart.ConsumerID,
		ProductID:  originItemCart.Product.ID,
		Count:      originItemCart.Count,
	}
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
	}
}

func MapToConsumerList(originList []pkgConsumerModel.Consumer) []Consumer {
	resultList := make([]Consumer, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToConsumer(v)
	}
	return resultList
}
