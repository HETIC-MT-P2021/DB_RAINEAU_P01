package models

import (
	"fmt"
	"gobdd/database"
)

type Employee struct {
	EmployeeNumber		string		`json:"employee_number"`
	LastName			string		`json:"last_name"`
	FirstName   		string		`json:"first_name"`
	Extension   		string		`json:"extension"`
	Email   			string		`json:"email"`
	ReportsTo			database.NullString		`json:"reports_to, omitempty"`
	JobTitle			string		`json:"job_title"`
	OfficeAddress		string		`json:"office_address"`
}

func(repository *Repository) GetEmployees() ([]*Employee, error){
	rows, err := repository.Conn.Query("SELECT e.employeeNumber, e.lastName, " + 
	"\ne.firstName, e.extension, e.email, e.reportsTo, e.jobTitle, o.addressLine1" +
	"\nAS officeAddress FROM employees e" + 
	"\nINNER JOIN offices o ON e.officeCode = o.officeCode")

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var employees []*Employee
	var employeeNumber, lastName, firstName, extension, email, jobTitle, officeAddress string
	var reportsTo database.NullString

	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&employeeNumber, &lastName, &firstName, &extension, &email, &reportsTo, &jobTitle, &officeAddress)

		if err != nil {
			return nil, fmt.Errorf("could not get employees : %v", err)
		}

		employee := &Employee{
			EmployeeNumber: employeeNumber,
			LastName:       lastName,
			FirstName:      firstName,
			Extension:      extension,
			Email:          email,
			ReportsTo:      reportsTo,
			JobTitle:       jobTitle,
			OfficeAddress:  officeAddress,
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return employees, nil

}

