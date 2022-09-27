package service

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/helper"
	"ark_backend_go/models/dto"
	repository "ark_backend_go/repositories"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type RemoteConfigServiceImpl struct {
	RemoteConfigRepository repository.RemoteConfigRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewRemoteConfigService(remoteConfigRepository repository.RemoteConfigRepository, DB *gorm.DB, validate *validator.Validate) RemoteConfigService {
	return &RemoteConfigServiceImpl{
		RemoteConfigRepository: remoteConfigRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

// GetRemoteConfig implements RemoteConfigService
func (service *RemoteConfigServiceImpl) GetRemoteConfig() *model.RemoteConfigGql {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return &model.RemoteConfigGql{}
	} else {
		remote := service.RemoteConfigRepository.GetRemoteConfig(tx)
		return dto.ToRemoteConfigResponseDTO(remote)
	}
}

// UpdateRemoteConfig implements RemoteConfigService
func (*RemoteConfigServiceImpl) UpdateRemoteConfig(request *model.RemoteConfigGql) model.RemoteConfigGql {
	panic("unimplemented")
}
