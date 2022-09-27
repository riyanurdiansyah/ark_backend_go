package graph

import (
	service "ark_backend_go/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PaymentService      service.PaymentService
	RemoteCongifService service.RemoteConfigService
}

func NewResolver(paymentService service.PaymentService, remoteConfigService service.RemoteConfigService) Resolver {
	return Resolver{
		PaymentService:      paymentService,
		RemoteCongifService: remoteConfigService,
	}
}
