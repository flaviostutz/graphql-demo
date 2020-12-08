package main

import (
	"net/http"
)

func main() {
	todo := Todo{}
	routes := makeRoutes(todo)
	http.ListenAndServe(":3000", routes)
}
