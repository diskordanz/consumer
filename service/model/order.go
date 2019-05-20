package model

import (
	"time"

	pkgOrderModel "github.com/diskordanz/web-consumer/pkg/order/model"
	pkgProductModel "github.com/diskordanz/web-consumer/pkg/product/model"
)

type Order struct {
	ID          uint64    `json:"id"`
	ConsumerID  uint64    `json:"consumer_id"`
	FranchiseID uint64    `json:"franchise_id"`
	Time        time.Time `json:"time"`
	Total       float32   `json:"total"`
	Status      string    `json:"status"`
}

type OrderItem struct {
	ID        uint64 `json:"id"`
	OrderID   uint64 `json:"order_id"`
	ProductID uint64 `json:"product_id"`
	Count     uint   `json:"count"`
}

type OrderItemWithProduct struct {
	ID      uint64  `json:"id"`
	OrderID uint64  `json:"order_id"`
	Product Product `json:"product"`
	Count   uint    `json:"count"`
}

type OrderWithItems struct {
	Order      Order                  `json:"order"`
	OrderItems []OrderItemWithProduct `json:"order_items"`
}

func MapToOrder(originOrder pkgOrderModel.Order) Order {
	return Order{
		ID:          originOrder.ID,
		ConsumerID:  originOrder.ConsumerID,
		FranchiseID: originOrder.FranchiseID,
		Time:        originOrder.Time,
		Total:       originOrder.Total,
		Status:      originOrder.Status,
	}
}

func MapToOrderWithItems(originOrder pkgOrderModel.Order, items []pkgOrderModel.OrderItem, products []pkgProductModel.Product) OrderWithItems {
	return OrderWithItems{
		Order:      MapToOrder(originOrder),
		OrderItems: MapToItemsWithProduct(items, products),
	}
}

func MapToItemsWithProduct(originList []pkgOrderModel.OrderItem, products []pkgProductModel.Product) []OrderItemWithProduct {
	resultList := make([]OrderItemWithProduct, len(originList), len(originList))
	for i, _ := range originList {
		resultList[i] = MapToOrderItemWithProduct(originList[i], products[i])
	}
	return resultList
}

func MapToItems(originList []pkgOrderModel.OrderItem) []OrderItem {
	resultList := make([]OrderItem, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToOrderItem(v)
	}
	return resultList
}
func MapToOrderItemWithProduct(originOrder pkgOrderModel.OrderItem, product pkgProductModel.Product) OrderItemWithProduct {
	return OrderItemWithProduct{
		ID:      originOrder.ID,
		OrderID: originOrder.OrderID,
		Count:   originOrder.Count,
		Product: MapToProduct(product),
	}
}

func MapToOrderItem(originOrder pkgOrderModel.OrderItem) OrderItem {
	return OrderItem{
		ID:        originOrder.ID,
		OrderID:   originOrder.OrderID,
		ProductID: originOrder.ProductID,
		Count:     originOrder.Count,
	}
}

func MapToOrderItemDB(originOrder OrderItem) pkgOrderModel.OrderItem {
	return pkgOrderModel.OrderItem{
		ID:        originOrder.ID,
		OrderID:   originOrder.OrderID,
		ProductID: originOrder.ProductID,
		Count:     originOrder.Count,
	}
}

func MapToOrderDB(originOrder Order) pkgOrderModel.Order {
	return pkgOrderModel.Order{
		ID:          originOrder.ID,
		ConsumerID:  originOrder.ConsumerID,
		FranchiseID: originOrder.FranchiseID,
		Time:        originOrder.Time,
		Total:       originOrder.Total,
		Status:      originOrder.Status,
	}
}

func MapToOrderList(originList []pkgOrderModel.Order) []Order {
	resultList := make([]Order, len(originList), len(originList))
	for i, v := range originList {
		resultList[i] = MapToOrder(v)
	}
	return resultList
}
