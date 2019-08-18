package models

//Employee Model is used to store the employee information
type Employee struct {
	EmployeeId   string `json:EmployeeId validation:NotNull`
	EmployeeName string `json:EmployeeName validation:NotNull`
	Address
}

//Address Model is used to store Address Information
type Address struct {
	AddressLine1 string `json:AddressLine1`
	AddressLine2 string `json:AddressLine2`
	City         string `json:City`
	State        string `json:State`
	Country      string `json:Country validation:NotNull`
}
