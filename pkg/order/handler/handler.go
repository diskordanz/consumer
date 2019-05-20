package handler

import (
	"github.com/diskordanz/web-consumer/pkg/order/model"
	"github.com/diskordanz/web-consumer/pkg/order/storage"
)

type OrderHandler struct {
	os storage.OrderStorage
}

func NewOrderHandler(os storage.OrderStorage) *OrderHandler {
	return &OrderHandler{os: os}
}

func (h *OrderHandler) ListOrders(id, count, offset int) ([]model.Order, error) {
	return h.os.List(id, count, offset)
}

func (h *OrderHandler) GetOrder(id int) (model.Order, error) {
	return h.os.Get(id)
}
func (h *OrderHandler) GetOrderItems(id int) ([]model.OrderItem, error) {
	return h.os.ListItems(id)
}

func (h *OrderHandler) CreateOrder(order model.Order) (model.Order, error) {
	return h.os.CreateOrder(order)
}

func (h *OrderHandler) CreateOrderItem(item model.OrderItem) (model.OrderItem, error) {
	return h.os.CreateOrderItem(item)
}
