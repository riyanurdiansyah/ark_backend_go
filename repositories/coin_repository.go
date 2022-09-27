package repository

import (
	"gorm.io/gorm"

	"ark_backend_go/models/entity"
)

type CoinRepository interface {
	GetCoinById(db *gorm.DB, userid int) (bool, *entity.Coin)
}
