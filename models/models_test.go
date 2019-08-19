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
