package handler

import (
	"github.com/diskordanz/web-consumer/pkg/product/model"
	"github.com/diskordanz/web-consumer/pkg/product/storage"
)

type ProductHandler struct {
	ps storage.ProductStorage
}

func NewProductHandler(ps storage.ProductStorage) *ProductHandler {
	return &ProductHandler{ps: ps}
}

func (h *ProductHandler) ListProducts(name string, count, offset int) ([]model.Product, error) {
	return h.ps.List(name, count, offset)
}

func (h *ProductHandler) ListProductsByCategory(id uint64, name string, count, offset int) ([]model.Product, error) {
	return h.ps.ListByCategory(id, name, count, offset)
}

func (h *ProductHandler) GetProductById(id int) (model.Product, error) {
	return h.ps.Get(id)
}
