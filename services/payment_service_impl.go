package service

import (
	"ark_backend_go/graph/model"
	"ark_backend_go/helper"
	"ark_backend_go/models/dto"
	"ark_backend_go/models/entity"
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

// AddPaymentMethod implements PaymentService
func (service *PaymentServiceImpl) AddPaymentMethod(request *model.PaymentMethodGql) model.PaymentMethodGql {
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		helper.PanicIfError(errorValidation)
		return *request
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		helper.PanicIfError(tx.Error)
		return *request
	} else {
		var status int
		if request.Status {
			status = 1
		} else {
			status = 0
		}
		payment := entity.PaymentMethod{
			ID:          request.ID,
			Value:       request.Value,
			Chanel:      request.Chanel,
			Code:        request.Code,
			Description: request.Description,
			Image:       request.Image,
			Limit:       request.Limit,
			Status:      status,
			Tipe:        request.Tipe,
			Title:       request.Title,
			TitleType:   request.TitleType,
		}

		response := service.PaymentRepository.AddPaymentMethod(tx, &payment)

		return *dto.ToPaymentMethodResponseDTO(response)
	}
}
