package repository

import (
	"gorm.io/gorm"

	"ark_backend_go/models/entity"
)

type PaymentRepository interface {
	GetPaymentMethod(db *gorm.DB) []*entity.PaymentMethod
}
