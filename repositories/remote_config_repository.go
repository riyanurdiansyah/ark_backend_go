package repository

import (
	"gorm.io/gorm"

	"ark_backend_go/models/entity"
)

type RemoteConfigRepository interface {
	GetRemoteConfig(db *gorm.DB) *entity.RemoteConfig
	AddPaymentMethod(db *gorm.DB, remote *entity.RemoteConfig) *entity.RemoteConfig
	UpdateRemoConfig(db *gorm.DB, remote *entity.RemoteConfig) *entity.RemoteConfig
}
