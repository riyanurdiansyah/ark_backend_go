package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ark_backend_go/graph/generated"
	"ark_backend_go/graph/model"
	"context"
	"fmt"
	"time"
)

// AddPaymentMethod is the resolver for the addPaymentMethod field.
func (r *mutationResolver) AddPaymentMethod(ctx context.Context, input model.NewPaymentMethod) (*model.PaymentMethodGql, error) {
	method := &model.PaymentMethodGql{
		ID:          input.ID,
		Value:       input.Value,
		Chanel:      input.Chanel,
		Code:        input.Chanel,
		Description: input.Description,
		Image:       input.Image,
		Limit:       input.Limit,
		Status:      input.Status,
		Tipe:        input.Tipe,
		Title:       input.Title,
		TitleType:   input.TitleType,
	}

	r.PaymentService.AddPaymentMethod(method)

	return method, nil
}

// Paymentmethod is the resolver for the paymentmethod field.
func (r *queryResolver) Paymentmethod(ctx context.Context) ([]*model.PaymentMethodGql, error) {
	data := r.PaymentService.GetPaymentMethod()
	return data, nil
}

// Method is the resolver for the method field.
func (r *subscriptionResolver) Method(ctx context.Context) (<-chan []*model.PaymentMethodGql, error) {
	ch := make(chan []*model.PaymentMethodGql)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			data := r.PaymentService.GetPaymentMethod()
			select {
			case ch <- data:
				// Our message went through, do nothing
			default:
				fmt.Println("Channel closed.")
				return
			}

		}
	}()
	return ch, nil
}

// Paymentmethod is the resolver for the paymentmethod field.
func (r *subscriptionResolver) Paymentmethod(ctx context.Context) (<-chan []*model.PaymentMethodGql, error) {
	ch := make(chan []*model.PaymentMethodGql)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			data := r.PaymentService.GetPaymentMethod()
			println("PING")
			select {
			case ch <- data:
				// Our message went through, do nothing
			default:
				fmt.Println("Channel closed.")
				return
			}

		}
	}()
	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
