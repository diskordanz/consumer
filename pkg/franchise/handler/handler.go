package handler

import (
	"github.com/diskordanz/web-consumer/pkg/franchise/model"
	"github.com/diskordanz/web-consumer/pkg/franchise/storage"
)

type FranchiseHandler struct {
	fs storage.FranchiseStorage
}

func NewFranchiseHandler(fs storage.FranchiseStorage) *FranchiseHandler {
	return &FranchiseHandler{fs: fs}
}

func (h *FranchiseHandler) ListFranchises(count, offset int) ([]model.Franchise, error) {
	return h.fs.List(count, offset)
}

func (h *FranchiseHandler) GetFranchise(id int) (model.Franchise, error) {
	return h.fs.Get(id)
}

/*
func (h *FranchiseHandler) SearchFranchisesByName(count, offset int, name string) ([]model.Franchise, error) {
	return h.fs.SearchFranchisesByName(count, offset, name)
}*/
