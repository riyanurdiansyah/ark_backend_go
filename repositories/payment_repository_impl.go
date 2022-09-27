package repository

import (
	"ark_backend_go/helper"
	"ark_backend_go/models/entity"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

// GetPaymentMethod implements PaymentRepository
func (*PaymentRepositoryImpl) GetPaymentMethod(db *gorm.DB) []*entity.PaymentMethod {
	var listPaymentMethod = []*entity.PaymentMethod{}
	result := db.Table("payment_method").Select("*").Scan(&listPaymentMethod)
	helper.PanicIfError(result.Error)

	return listPaymentMethod
}

// AddPaymentMethod implements PaymentRepository
func (*PaymentRepositoryImpl) AddPaymentMethod(db *gorm.DB, payment *entity.PaymentMethod) *entity.PaymentMethod {
	result := db.Table("payment_method").Select("*").Create(&payment)
	if result.Error != nil {
		///handle panic
		panic(result.Error)
	}
	return payment
}
