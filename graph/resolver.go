package graph

import (
	service "ark_backend_go/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PaymentService service.PaymentService
}

func NewResolver(paymentService service.PaymentService) Resolver {
	return Resolver{
		PaymentService: paymentService,
	}
}
