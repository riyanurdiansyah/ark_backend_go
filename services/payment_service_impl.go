package service

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/helper"
	"ark_backend_go/models/dto"
	repository "ark_backend_go/repositories"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, DB *gorm.DB, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

// GetPaymentMethod implements PaymentService
func (service *PaymentServiceImpl) GetPaymentMethod() []*model.PaymentMethodGql {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return []*model.PaymentMethodGql{}
	} else {
		listPaymentMethod := service.PaymentRepository.GetPaymentMethod(tx)
		return dto.ToListPaymentMethodResponseGQL(listPaymentMethod)
	}
}
