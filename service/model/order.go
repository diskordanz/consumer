package model

import (
	pkgOrderModel "github.com/diskordanz/consumer/pkg/order/model"
)

type Order struct {
	ID         uint64        `json:"id"`
	ConsumerID uint64        `json:"consumer_id"`
	Products   map[int64]int `json:"products"`
	Date       string        `json:"date"`
	Total      float32       `json:"total"`
	Status     string        `json:"status"`
}

func MapToOrder(originOrder pkgOrderModel.Order) Order {
	return Order{
		ID:         originOrder.ID,
		ConsumerID: originOrder.ConsumerID,
		Products:   originOrder.Products,
		Date:       originOrder.Date,
		Total:      originOrder.Total,
		Status:     originOrder.Status,
	}
}

func MapToOrderDB(originOrder Order) pkgOrderModel.Order {
	return pkgOrderModel.Order{
		ID:         originOrder.ID,
		ConsumerID: originOrder.ConsumerID,
		Products:   originOrder.Products,
		Date:       originOrder.Date,
		Total:      originOrder.Total,
		Status:     originOrder.Status,
	}
}

func MapToOrderList(originList []pkgOrderModel.Order) []Order {
	resultList := make([]Order, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToOrder(v)
	}
	return resultList
}
