package service

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/models/dto"
	repository "ark_backend_go/repositories"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type CoinServiceImpl struct {
	CoinRepository repository.CoinRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCoinService(coinRepository repository.CoinRepository, DB *gorm.DB, validate *validator.Validate) CoinService {
	return &CoinServiceImpl{
		CoinRepository: coinRepository,
		DB:             DB,
		Validate:       validate,
	}
}

// GetCoinById implements CoinService
func (service *CoinServiceImpl) GetCoinById(userid int) (bool, *model.CoinGql) {
	tx := service.DB.Begin()
	if tx.Error != nil {
		return true, nil
	} else {
		status, data := service.CoinRepository.GetCoinById(tx, userid)
		return status, dto.ToCoinResponseDTO(data)
	}
}
