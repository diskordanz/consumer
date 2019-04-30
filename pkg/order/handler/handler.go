package handler

import (
	"github.com/diskordanz/consumer/pkg/order/model"
	"github.com/diskordanz/consumer/pkg/order/storage"
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

func (h *OrderHandler) GetOrder(consumerID, orderID int) (model.Order, error) {
	return h.os.Get(consumerID, orderID)
}
