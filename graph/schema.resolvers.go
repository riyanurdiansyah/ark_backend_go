package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ark_backend_go/graph/generated"
	"ark_backend_go/graph/model"
	"context"
)

// Paymentmethod is the resolver for the paymentmethod field.
func (r *queryResolver) Paymentmethod(ctx context.Context) (*model.PaymentMethodResponse, error) {
	data := r.PaymentService.GetPaymentMethod()
	response := &model.PaymentMethodResponse{
		Code:   200,
		Status: true,
		Data:   data,
	}

	return response, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
