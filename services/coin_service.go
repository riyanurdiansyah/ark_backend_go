package service

import (
	"ark_backend_go/graph/model"
)

type CoinService interface {
	GetCoinById(userid int) (bool, *model.CoinGql)
}
