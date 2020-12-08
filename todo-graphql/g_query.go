package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
		"user": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"order": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
func buildQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{

			/*
				curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
			*/
			"todo": &graphql.Field{
				Type:        todoType,
				Description: "Get single todo",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tid, _ := params.Args["id"].(string)
					response, err := sendJSONRequest("GET", fmt.Sprintf("%s/todos/%s", opt.TodoRestURL, tid), nil)
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

			/*
				curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
			*/
			"todoList": &graphql.Field{
				Type:        graphql.NewList(todoType),
				Description: "List of todos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					response, err := sendJSONRequest("GET", fmt.Sprintf("%s/todos", opt.TodoRestURL), nil)
					if err != nil {
						return nil, err
					}
					resdata, err := processResponse(response, 200, true)
					if err != nil {
						return nil, err
					}
					return resdata, nil
				},
			},
		},
	})
}
