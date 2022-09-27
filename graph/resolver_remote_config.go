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

// UpdateRemoteConfig is the resolver for the updateRemoteConfig field.
func (r *mutationResolver) UpdateRemoteConfig(ctx context.Context, input *model.InputRemoteConfigGql) (*model.RemoteConfigGql, error) {
	panic(fmt.Errorf("not implemented: UpdateRemoteConfig - updateRemoteConfig"))
}

// Remoteconfig is the resolver for the remoteconfig field.
func (r *subscriptionResolver) Remoteconfig(ctx context.Context) (<-chan *model.RemoteConfigGql, error) {
	ch := make(chan *model.RemoteConfigGql)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			data := r.RemoteCongifService.GetRemoteConfig()
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
func (r *Resolver) RemoteConfigMutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) RemoteConfigQuery() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) RemoteConfigSubscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}
