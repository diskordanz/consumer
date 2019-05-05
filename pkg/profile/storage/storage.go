package storage

import (
	"github.com/diskordanz/consumer/pkg/profile/model"
)

type ProfileStorage interface {
	// Get list of size *count* of Profiles starting from *offset*
	Get(int) (model.Consumer, error)
}
