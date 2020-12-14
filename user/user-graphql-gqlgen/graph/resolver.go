package graph

import "github.com/flaviostutz/graphql-demo/user-graphql-gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var users []*model.User

func init() {
	users = make([]*model.User, 0)
}
