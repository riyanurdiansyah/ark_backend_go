package repository

import (
	"ark_backend_go/helper"
	"ark_backend_go/models/entity"

	"gorm.io/gorm"
)

type RemoteConfigRepositoryImpl struct {
}

func NewRemoteConfigRepository() RemoteConfigRepository {
	return &RemoteConfigRepositoryImpl{}
}

// AddPaymentMethod implements RemoteConfigRepository
func (*RemoteConfigRepositoryImpl) AddPaymentMethod(db *gorm.DB, remote *entity.RemoteConfig) *entity.RemoteConfig {
	panic("unimplemented")
}

// GetRemoteConfig implements RemoteConfigRepository
func (*RemoteConfigRepositoryImpl) GetRemoteConfig(db *gorm.DB) *entity.RemoteConfig {
	var remoteConfig = []*entity.RemoteConfig{}
	result := db.Table("remote_config").Select("*").Scan(&remoteConfig)
	helper.PanicIfError(result.Error)

	return remoteConfig[0]
}

// UpdateRemoConfig implements RemoteConfigRepository
func (*RemoteConfigRepositoryImpl) UpdateRemoConfig(db *gorm.DB, remote *entity.RemoteConfig) *entity.RemoteConfig {
	result :=
		db.Table("remote_config").Select("*").Updates(&remote)
	helper.PanicIfError(result.Error)
	return remote
}
