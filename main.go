package main

import (
	"employee-service/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	println("Server is gonna start ")

}

// Create a new route and listen on post 8080 for all the requests
func main() {
	route := mux.NewRouter()
	handlers.HandleFunction(route)
	err := http.ListenAndServe(":8081", route)
	if err != nil {
		println("Error in running the server ", err.Error())
		panic(err)
	}
}
