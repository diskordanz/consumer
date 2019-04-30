package handler

import (
	"github.com/diskordanz/consumer/pkg/category/model"
	"github.com/diskordanz/consumer/pkg/category/storage"
)

type CategoryHandler struct {
	fs storage.CategoryStorage
}

func NewCategoryHandler(fs storage.CategoryStorage) *CategoryHandler {
	return &CategoryHandler{fs: fs}
}

func (h *CategoryHandler) ListCategories() ([]model.Category, error) {
	return h.fs.List()
}
