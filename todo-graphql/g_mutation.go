package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// root mutation
func buildMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
			*/
			"createTodo": &graphql.Field{
				Type:        todoType, // the return type for this field
				Description: "Create new todo",
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"user": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					response, err := sendJSONRequest("POST", fmt.Sprintf("%s/todos", opt.TodoRestURL), params.Args)
					if err != nil {
						return nil, err
					}
					resdata, err := processResponse(response, 201, false)
					if err != nil {
						return nil, err
					}
					return resdata, nil
				},
			},
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
			*/
			"updateTodo": &graphql.Field{
				Type:        graphql.String,
				Description: "Update existing todo",
				Args: graphql.FieldConfigArgument{
					"done": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"user": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tid, _ := params.Args["id"].(string)

					response, err := sendJSONRequest("PATCH", fmt.Sprintf("%s/todos/%s", opt.TodoRestURL, tid), params.Args)
					if err != nil {
						return nil, err
					}
					resdata, err := processResponse(response, 200, false)
					if err != nil {
						return nil, err
					}
					return resdata, nil
				},
			},
			"deleteTodo": &graphql.Field{
				Type:        graphql.String,
				Description: "Delete existing todo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tid, _ := params.Args["id"].(string)

					response, err := sendJSONRequest("DELETE", fmt.Sprintf("%s/todos/%s", opt.TodoRestURL, tid), params.Args)
					if err != nil {
						return nil, err
					}
					resdata, err := processResponse(response, 200, false)
					if err != nil {
						return nil, err
					}
					return resdata, nil
				},
			},
		},
	})
}
