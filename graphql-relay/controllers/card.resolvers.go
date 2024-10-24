package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/Reimei1213/lab/graphql-relay/interface/presenter"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"
)

// AssignedUser is the resolver for the assignedUser field.
func (r *cardResolver) AssignedUser(ctx context.Context, obj *model.Card) (*model.User, error) {
	if obj.AssignedUser.ID == "" {
		return nil, nil
	}

	u, err := r.UserInputport.Get(ctx, obj.AssignedUser.ID)
	if err != nil {
		return nil, err
	}
	return presenter.NewUser(u).ToGraphqlModel(), nil
}

// AddCard is the resolver for the addCard field.
func (r *mutationResolver) AddCard(ctx context.Context, input model.AddCardInput) (*model.AddCardPayload, error) {
	if input.UserID != nil {
		_, userID, err := presenter.DecodeGraphqlID(*input.UserID)
		if err != nil {
			return nil, err
		}
		input.UserID = &userID
	}
	id, err := r.CardInputport.Create(ctx, &input)
	if err != nil {
		return nil, err
	}
	return &model.AddCardPayload{ID: presenter.EncodeGraphqlID(presenter.NodeTypeCard, id)}, nil
}

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context, first *int, after *string, last *int, before *string) (*model.Connection, error) {
	panic(fmt.Errorf("not implemented: Cards - cards"))
}

// Card returns graph.CardResolver implementation.
func (r *Resolver) Card() graph.CardResolver { return &cardResolver{r} }

type cardResolver struct{ *Resolver }