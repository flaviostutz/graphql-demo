package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/flaviostutz/graphql-demo/todo-graphql-gqlgen/graph/generated"
	"github.com/flaviostutz/graphql-demo/todo-graphql-gqlgen/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, title *string, text *string, user *string, order *int) (*model.Todo, error) {
	t := model.Todo{
		Title: title,
		Text:  text,
		User:  user,
		Order: order,
	}
	response, err := sendJSONRequest("POST", fmt.Sprintf("%s/todos", todoRestURL), t)
	if err != nil {
		return nil, err
	}
	todo, err := processResponseTodo(response, 201)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*string, error) {
	response, err := sendJSONRequest("DELETE", fmt.Sprintf("%s/todos/%s", todoRestURL, id), nil)
	if err != nil {
		return nil, err
	}
	resdata, err := processResponseText(response, 200)
	if err != nil {
		return nil, err
	}
	return &resdata, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, text *string, order *int, id string, done *bool, user *string, title *string) (*string, error) {
	t := model.Todo{
		Title: title,
		Text:  text,
		User:  user,
		Order: order,
		Done:  done,
	}
	response, err := sendJSONRequest("POST", fmt.Sprintf("%s/todos/%s", todoRestURL, id), t)
	if err != nil {
		return nil, err
	}
	txt, err := processResponseText(response, 201)
	if err != nil {
		return nil, err
	}
	return &txt, nil
}

func (r *queryResolver) Todo(ctx context.Context, id *string) (*model.Todo, error) {
	response, err := sendJSONRequest("GET", fmt.Sprintf("%s/todos/%s", todoRestURL, *id), nil)
	if err != nil {
		return nil, err
	}
	resdata, err := processResponseTodo(response, 200)
	if err != nil {
		return nil, err
	}
	return resdata, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	response, err := sendJSONRequest("GET", fmt.Sprintf("%s/todos", todoRestURL), nil)
	if err != nil {
		return nil, err
	}
	resdata, err := processResponseList(response, 200)
	if err != nil {
		return nil, err
	}
	return resdata, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
