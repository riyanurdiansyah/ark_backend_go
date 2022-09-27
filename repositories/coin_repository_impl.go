package repository

import (
	"ark_backend_go/helper"
	"ark_backend_go/models/entity"

	"gorm.io/gorm"
)

type CoinRepositoryImpl struct {
}

func NewCoinRepository() CoinRepository {
	return &CoinRepositoryImpl{}
}

// GetCoinById implements CoinRepository
func (*CoinRepositoryImpl) GetCoinById(db *gorm.DB, userid int) (bool, *entity.Coin) {
	var coin = entity.Coin{}
	var total int64
	result := db.Table("coins").Select("*").Where("user_id = ?", userid).Scan(&coin).Count(&total)
	helper.PanicIfError(result.Error)
	if total > 0 {
		return true, &coin
	}
	return false, &coin
}
