package entity

type RemoteConfig struct {
	BaseUrl               string `gorm:"column:base_url"`
	BaseURLPrakerja       string `gorm:"column:base_url_prakerja"`
	ForceLogout           int    `gorm:"column:force_logout"`
	IsChecking            int    `gorm:"column:is_checking"`
	IsMaintenance         int    `gorm:"column:is_maintenance"`
	IsMaintenancePrakerja int    `gorm:"column:is_maintenance_prakerja"`
	NewVersion            int    `gorm:"column:new_version"`
}
