package repository

import (
	"gorm.io/gorm"

	"ark_backend_go/models/entity"
)

type PaymentRepository interface {
	GetPaymentMethod(db *gorm.DB) []*entity.PaymentMethod
	AddPaymentMethod(db *gorm.DB, payment *entity.PaymentMethod) *entity.PaymentMethod
}
