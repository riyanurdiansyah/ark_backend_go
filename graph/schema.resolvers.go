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
func (r *mutationResolver) AddPaymentMethod(ctx context.Context, input model.InputNewPaymentGql) (*model.PaymentMethodGql, error) {
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

// UpdateRemoteConfig is the resolver for the updateRemoteConfig field.
func (r *mutationResolver) UpdateRemoteConfig(ctx context.Context, input model.InputRemoteConfigGql) (*model.RemoteConfigGql, error) {
	data := r.RemoteConfigService.UpdateRemoteConfig(&input)

	return &data, nil
}

// UpdateCoin is the resolver for the update_coin field.
func (r *mutationResolver) UpdateCoin(ctx context.Context, input *model.InputCoinGql) (*model.CoinGql, error) {
	panic(fmt.Errorf("not implemented: UpdateCoin - update_coin"))
}

// PaymentMethod is the resolver for the payment_method field.
func (r *queryResolver) PaymentMethod(ctx context.Context) ([]*model.PaymentMethodGql, error) {
	data := r.PaymentService.GetPaymentMethod()
	return data, nil
}

// RemoteConfig is the resolver for the remote_config field.
func (r *queryResolver) RemoteConfig(ctx context.Context) (*model.RemoteConfigGql, error) {
	data := r.RemoteConfigService.GetRemoteConfig()
	return data, nil
}

// PaymentMethod is the resolver for the payment_method field.
func (r *subscriptionResolver) PaymentMethod(ctx context.Context) (<-chan []*model.PaymentMethodGql, error) {
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

// RemoteConfig is the resolver for the remote_config field.
func (r *subscriptionResolver) RemoteConfig(ctx context.Context) (<-chan *model.RemoteConfigGql, error) {
	ch := make(chan *model.RemoteConfigGql)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			data := r.RemoteConfigService.GetRemoteConfig()
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

// Coin is the resolver for the coin field.
func (r *subscriptionResolver) Coin(ctx context.Context, userid int) (<-chan *model.CoinGql, error) {
	ch := make(chan *model.CoinGql)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			status, data := r.CoinService.GetCoinById(userid)
			println("FROM SERVER =========> GET COIN")
			if status {
				select {
				case ch <- data:
					// Our message went through, do nothing
				default:
					fmt.Println("Channel closed.")
					return
				}
			} else {
				select {
				case ch <- &model.CoinGql{}:
					// Our message went through, do nothing
				default:
					fmt.Println("Channel closed.")
					return
				}
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
