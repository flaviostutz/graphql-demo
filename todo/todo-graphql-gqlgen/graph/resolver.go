package graph

import (
	"fmt"
	"os"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var todoRestURL = os.Getenv("TODO_SERVICE_URL")

func init() {
	if todoRestURL == "" {
		fmt.Println("env TODO_SERVICE_URL is required")
		os.Exit(1)
	}
}
