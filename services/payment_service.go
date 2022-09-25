package service

import (
	"ark_backend_go/graph/model"
)

type PaymentService interface {
	GetPaymentMethod() []*model.PaymentMethodGql
}
