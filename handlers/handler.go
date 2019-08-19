package handlers

import (
	"employee-service/models"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

//Users credentials
var usersList = map[string]string{
	"Gowtham": "12345", //R293dGhhbToxMjM0NQ==
}

//define the handler functions
func HandleFunction(route *mux.Router) {
	route.Handle("/Employee/{id}", authenticate(EmployeeHandler)).Methods("GET", "PUT", "DELETE")
	route.Handle("/Employee", authenticate(EmployeeHandler))

}

//handler will authenticate the user based request header and if success pass to the EmployeeHandler
//otherwise return to the user with response as 401 UNAuthorized
func authenticate(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if validateUser(request) {
			request.Header.Set("ValidUser", "YES") // JUST FOR TESTING DEMO
			next.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
		}
		return

	})
}

//validate the user credential which is form of base64 encoded
func validateUser(request *http.Request) bool {
	authValue := request.Header.Get("Authorization") // Basic <value>
	if authValue == "" {
		return false
	}
	decodedValue, err := base64.StdEncoding.DecodeString(strings.Split(authValue, " ")[1])
	if err != nil {
		return false
	}
	data := strings.Split(string(decodedValue), ":")
	if len(data) != 2 {
		return false
	}
	if password, ok := usersList[data[0]]; ok {
		if password == data[1] {
			return true
		}
	}

	return false
}

//EmployeeHandler will be used to perform the CRUD operations for the /Employee API endpoint
func EmployeeHandler(writer http.ResponseWriter, request *http.Request) {
	methodType := request.Method
	switch methodType {
	case "GET":
		emp := models.Employee{
			EmployeeId:   "1",
			EmployeeName: "Gowtham",
			Address:      models.Address{},
		}
		response, err := json.Marshal(emp)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.Write([]byte(response))
	case "POST":
		emp := models.Employee{}
		err := json.NewDecoder(request.Body).Decode(&emp)
		if err != nil {
			writer.WriteHeader(http.StatusCreated)
			writer.Write([]byte("Created Successfully"))
		}

	case "PUT":
		pathParameters := mux.Vars(request)
		employeeId := pathParameters["id"]
		writer.Write([]byte("the employee information for the employee with id " + employeeId + " is udpdated"))
	case "DELETE":

	default:

	}
}
