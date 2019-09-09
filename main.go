package main

import (
	"employee-service/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func init() {
	println("Server is gonna start ")

}

// Create a new route and listen on post 8080 for all the requests
func main() {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "9090"
	}
	route := mux.NewRouter()
	handlers.HandleFunction(route)
	err := http.ListenAndServe(":"+port, route)
	if err != nil {
		println("Error in running the server ", err.Error())
		panic(err)
	}
}
