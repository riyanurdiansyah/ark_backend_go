package dto

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/models/entity"
)

func ToCoinResponseDTO(coin *entity.Coin) *model.CoinGql {
	var (
		isOldUser   bool
		isCompleted bool
	)

	if coin.IsOldUser == 1 {
		isOldUser = true
	} else {
		isOldUser = false
	}

	if coin.IsCompleted == 1 {
		isCompleted = true
	} else {
		isCompleted = false
	}
	return &model.CoinGql{
		UserID:      coin.UserID,
		IsOldUser:   isOldUser,
		IsCompleted: isCompleted,
		CreatedAt:   coin.CreatedAt,
		UpdatedAt:   coin.UpdatedAt,
		Coins:       coin.Coin,
	}
}
