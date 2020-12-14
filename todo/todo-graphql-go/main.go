package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
)

type Opt struct {
	TodoRestURL string
}

var opt Opt

func main() {
	opt.TodoRestURL = os.Getenv("TODO_SERVICE_URL")
	if opt.TodoRestURL == "" {
		fmt.Println("env TODO_SERVICE_URL is required")
		os.Exit(1)
	}

	schema, err := buildTodoSchema()
	if err != nil {
		fmt.Printf("Error building schema. err=%s\n", err)
		os.Exit(2)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/", h)

	// Display some basic instructions
	fmt.Println("Now server is running on port 4000")
	fmt.Println("Get single todo: curl -g 'http://localhost:4000/graphql?query={todo(id:\"b\"){id,title,done}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:4000/graphql?query=mutation+_{createTodo(title:\"My+new+todo\"){id,title,done}}'")
	fmt.Println("Update todo: curl -g 'http://localhost:4000/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:4000/graphql?query={todoList{id,title,done}}'")

	http.ListenAndServe(":4000", nil)
}
