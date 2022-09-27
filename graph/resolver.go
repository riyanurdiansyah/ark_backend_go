package graph

import (
	service "ark_backend_go/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PaymentService      service.PaymentService
	RemoteConfigService service.RemoteConfigService
	CoinService         service.CoinService
}

func NewResolver(paymentService service.PaymentService, remoteConfigService service.RemoteConfigService, coinService service.CoinService) Resolver {
	return Resolver{
		PaymentService:      paymentService,
		RemoteConfigService: remoteConfigService,
		CoinService:         coinService,
	}
}
