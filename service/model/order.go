package model

import (
	pkgOrderModel "github.com/diskordanz/consumer/pkg/order/model"
)

type Order struct {
	ID        uint64 `json:"id"`
	CountryID int32  `json:"country_id"`
	Name      string `json:"name"`
}

func MapToOrder(originOrder pkgOrderModel.Order) Order {
	return Order{ID: originOrder.ID}
}

func MapToOrderList(originList []pkgOrderModel.Order) []Order {
	resultList := make([]Order, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToOrder(v)
	}
	return resultList
}
