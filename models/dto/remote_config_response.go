package dto

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/models/entity"
)

func ToRemoteConfigEntity(remote *model.InputRemoteConfigGql) *entity.RemoteConfig {
	var (
		forceLogout           int
		isChecking            int
		isImportant           int
		isMaintenance         int
		isMaintenancePrakerja int
	)

	if remote.ForceLogout {
		forceLogout = 1
	} else {
		forceLogout = 0
	}

	if remote.IsChecking {
		isChecking = 1
	} else {
		isChecking = 0
	}
	if remote.IsImportant {
		isImportant = 1
	} else {
		isImportant = 0
	}

	if remote.IsMaintenance {
		isMaintenance = 1
	} else {
		isMaintenance = 0
	}

	if remote.IsMaintenancePrakerja {
		isMaintenancePrakerja = 1
	} else {
		isMaintenancePrakerja = 0
	}

	return &entity.RemoteConfig{
		ID:                    1,
		BaseUrl:               remote.BaseURL,
		BaseURLPrakerja:       remote.BaseURLPrakerja,
		ForceLogout:           forceLogout,
		IsChecking:            isChecking,
		IsImportant:           isImportant,
		IsMaintenance:         isMaintenance,
		IsMaintenancePrakerja: isMaintenancePrakerja,
		NewVersion:            remote.NewVersion,
	}
}

func ToRemoteConfigResponseDTO(remote *entity.RemoteConfig) *model.RemoteConfigGql {
	var (
		forceLogout           bool
		isChecking            bool
		isImportant           bool
		isMaintenance         bool
		isMaintenancePrakerja bool
	)

	if remote.ForceLogout == 1 {
		forceLogout = true
	} else {
		forceLogout = false
	}

	if remote.IsChecking == 1 {
		isChecking = true
	} else {
		isChecking = false
	}

	if remote.IsImportant == 1 {
		isImportant = true
	} else {
		isImportant = false
	}

	if remote.IsMaintenance == 1 {
		isMaintenance = true
	} else {
		isMaintenance = false
	}

	if remote.IsMaintenancePrakerja == 1 {
		isMaintenancePrakerja = true
	} else {
		isMaintenancePrakerja = false
	}

	return &model.RemoteConfigGql{
		BaseURL:              &remote.BaseUrl,
		BaseURLPrakerja:      &remote.BaseURLPrakerja,
		ForceLogout:          &forceLogout,
		IsChecking:           &isChecking,
		IsImportant:          &isImportant,
		IsMantenance:         &isMaintenance,
		IsMantenancePrakerja: &isMaintenancePrakerja,
		NewVersion:           &remote.NewVersion,
	}
}
