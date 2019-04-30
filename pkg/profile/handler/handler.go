package handler

import (
	"github.com/diskordanz/consumer/pkg/profile/model"
	"github.com/diskordanz/consumer/pkg/profile/storage"
)

type ProfileHandler struct {
	fs storage.ProfileStorage
}

func NewProfileHandler(fs storage.ProfileStorage) *ProfileHandler {
	return &ProfileHandler{fs: fs}
}

func (h *ProfileHandler) GetProfile(id int) (model.Profile, error) {
	return h.fs.Get(id)
}
