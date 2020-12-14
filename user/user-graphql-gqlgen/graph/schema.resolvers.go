package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/flaviostutz/graphql-demo/user-graphql-gqlgen/graph/generated"
	"github.com/flaviostutz/graphql-demo/user-graphql-gqlgen/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, name string, bio *string) (*model.User, error) {
	m := &model.User{Name: name, Bio: bio, ID: fmt.Sprintf("%d", time.Now().Nanosecond())}
	users = append(users, m)
	return m, nil
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	for _, u := range users {
		if u.ID == *id {
			return u, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
