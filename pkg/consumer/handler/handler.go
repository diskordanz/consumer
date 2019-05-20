package handler

import (
	"github.com/diskordanz/web-consumer/pkg/consumer/model"
	"github.com/diskordanz/web-consumer/pkg/consumer/storage"
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

func (h *ConsumerHandler) CreateCartItem(item model.CartItem) (model.CartItem, error) {
	return h.us.CreateCartItem(item)
}

func (h *ConsumerHandler) UpdateCartItem(item model.CartItem) (model.CartItem, error) {
	return h.us.UpdateCartItem(item)
}

func (h *ConsumerHandler) DeleteCartItem(item model.CartItem) error {
	return h.us.DeleteCartItem(item)
}

func (h *ConsumerHandler) GetCartItem(id, product_id int) (model.CartItem, error) {
	return h.us.GetCartItem(id, product_id)
}
