package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//test method name should start with Test<methodName- start with caps for even private method>
//demo on testing the handler directly
func TestEmployeeHandler(t *testing.T) {
	//create a new request
	req, err := http.NewRequest("GET", "/Employee/1", nil)
	if err != nil {
		log.Fatal("Error in Creating the new request") //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	//to record the response
	responseRecord := httptest.NewRecorder()
	handler := http.HandlerFunc(EmployeeHandler)
	handler.ServeHTTP(responseRecord, req) // call the handler

	//check the response and raise an error by using the log.Fatal or t.Error or t.Fail or t.Errorf
	if responseRecord.Body == nil {
		t.Error("Response is null and exptected is one entry") //equal to log followed by fail
		//print like : handler_test.go:24: Response is null and exptected is one entry
	}
	if responseRecord.Body.String() != `{"EmployeeId":"1","EmployeeName":"Gowtham","AddressLine1":"","AddressLine2":"","City":"","State":"","Country":""}` {
		t.Fail() // fail the test case but continue execution
	}

}

//demo on testing the middleware
func TestAuthenticate(t *testing.T) {
	//create a new request
	req, err := http.NewRequest("GET", "/Employee/1", nil)
	if err != nil {
		log.Fatal("Error in Creating the new request") //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	//to record the response
	responseRecord := httptest.NewRecorder()
	//if any data is set in authentication handler , that we can check here
	empHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
	handler := http.Handler(authenticate(empHandler))
	handler.ServeHTTP(responseRecord, req) // call the handler

	if responseRecord.Code != 401 {
		t.Error(" expected 401 ")
	}

}

//demo on testing the middleware with authentication and further handler request data is tested
func TestAuthenticate_2(t *testing.T) {
	//create a new request
	req, err := http.NewRequest("GET", "/Employee/1", nil)
	req.SetBasicAuth("Gowtham", "12345")
	if err != nil {
		log.Fatal("Error in Creating the new request") //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	//to record the response
	responseRecord := httptest.NewRecorder()
	//if any data is set in authentication handler , that we can check here
	empHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// in this section we can test request parameter coming to the handler
		if req.Header.Get("ValidUser") != "YES" {
			t.Error("expected as valid user")
		}
	})
	handler := http.Handler(authenticate(empHandler)) // here our part employee handler mock will be executed instead of original mock
	handler.ServeHTTP(responseRecord, req)            // call the handler

}
