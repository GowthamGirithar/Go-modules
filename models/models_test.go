package models_test

import (
	"employee-service/models" //since package is different import is required and here we cannot test private methods
	"github.com/stretchr/testify/assert"
	"testing"
)

//here we have used the external test file concept
func TestEmployeeModel(t *testing.T) {
	emp := models.Employee{
		EmployeeId:   "123",
		EmployeeName: "",
		Address:      models.Address{},
	}
	//we can directly check or use the thrid party packages to test the models assert
	assert.Equal(t, "123", emp.EmployeeId)
}

//here we will be testing many scenarios in one method
func TestEmployeeModel2(t *testing.T) {
	emp := models.Employee{
		EmployeeId:   "123",
		EmployeeName: "GG",
		Address:      models.Address{},
	}

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // tells that this is helper method and when it fails, it will say in detail which test case failed
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		expected := "123"
		actual := emp.EmployeeId
		assertCorrectMessage(t, actual, expected)
	})
	t.Run("Employee Name check", func(t *testing.T) {
		expected := "GG"
		actual := emp.EmployeeName
		assertCorrectMessage(t, actual, expected)
	})
}
