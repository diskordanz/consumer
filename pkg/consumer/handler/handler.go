package handler

import (
	"github.com/diskordanz/consumer/pkg/consumer/model"
	"github.com/diskordanz/consumer/pkg/consumer/storage"
)

type ConsumerHandler struct {
	us storage.ConsumerStorage
}

func NewConsumerHandler(us storage.ConsumerStorage) *ConsumerHandler {
	return &ConsumerHandler{us: us}
}

func (h *ConsumerHandler) GetConsumer(id int) (model.Consumer, error) {
	return h.us.Get(id)
}

func (h *ConsumerHandler) CreateConsumer(consumer model.Consumer) (model.Consumer, error) {
	return h.us.Create(consumer)
}

func (h *ConsumerHandler) UpdateConsumer(consumer model.Consumer) (model.Consumer, error) {
	return h.us.Update(consumer)
}

func (h *ConsumerHandler) GetCart(id, count, offset int) ([]model.CartItem, error) {
	return h.us.Cart(id, count, offset)
}
