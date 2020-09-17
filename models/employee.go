package models

type Employee struct {
	EmployeeNumber		string		`json:"employeeNumber"`
	LastName			string		`json:"lastName"`
	FirstName   		string		`json:"firstName"`
	Extension   		string		`json:"extension"`
	Email   			string		`json:"email"`
	Office				Office		`json:"_"`
	ReportsTo			*Employee	`json:"_,omitempty"`
	JobTitle			string		`json:"jobsTitle"`
}


