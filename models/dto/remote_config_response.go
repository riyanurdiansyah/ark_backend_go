package dto

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/models/entity"
)

func ToRemoteConfigResponseDTO(remote *entity.RemoteConfig) *model.RemoteConfigGql {

	return &model.RemoteConfigGql{
		BaseURL:              &remote.BaseUrl,
		BaseURLPrakerja:      &remote.BaseURLPrakerja,
		ForceLogout:          &remote.ForceLogout,
		IsChecking:           &remote.IsChecking,
		IsMantenance:         &remote.IsMaintenance,
		IsMantenancePrakerja: &remote.IsMaintenancePrakerja,
		NewVersion:           &remote.NewVersion,
	}
}
