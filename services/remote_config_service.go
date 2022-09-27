package service

import (
	"ark_backend_go/graph/model"
)

type RemoteConfigService interface {
	GetRemoteConfig() *model.RemoteConfigGql
	UpdateRemoteConfig(request *model.RemoteConfigGql) model.RemoteConfigGql
}
