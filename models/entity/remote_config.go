package entity

type RemoteConfig struct {
	ID                    int    `gorm:"column:id;primaryKey;autoIncrement"`
	BaseUrl               string `gorm:"column:base_url"`
	ForceLogout           int    `gorm:"column:force_logout"`
	IsChecking            int    `gorm:"column:is_checking"`
	IsImportant           int    `gorm:"column:is_important"`
	IsMaintenance         int    `gorm:"column:is_maintenance"`
	IsMaintenancePrakerja int    `gorm:"column:is_maintenance_prakerja"`
	NewVersion            int    `gorm:"column:new_version"`
	BaseURLPrakerja       string `gorm:"column:base_url_prakerja"`
}
