package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Create a new route and listen on post 8080 for all the requests
func main() {
	route := mux.NewRouter()
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		println("Error in running the server ", err.Error())
		panic(err)
	}
}
