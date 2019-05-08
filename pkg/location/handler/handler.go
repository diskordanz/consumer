package handler

import (
	"github.com/diskordanz/consumer/pkg/location/model"
	"github.com/diskordanz/consumer/pkg/location/storage"
)

type LocationHandler struct {
	ls storage.LocationStorage
}

func NewLocationHandler(ls storage.LocationStorage) *LocationHandler {
	return &LocationHandler{ls: ls}
}

func (h *LocationHandler) GetLocationsOfFranchise(id, count, offset int) ([]model.Location, error) {
	return h.ls.List(id, count, offset)
}

/*
func (h *LocationHandler) GetLocationsOfFranchiseByName(id, count, offset int, name string) ([]model.Location, error) {
	return h.ls.ListOfLocationsByName(id, count, offset, name)
}*/
