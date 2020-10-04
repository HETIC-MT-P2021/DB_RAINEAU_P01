package models

import (
	"database/sql"
	"fmt"
	"gobdd/database"
)

type Office struct {
	OfficeCode		string	`json:"office_code"`
	City			string	`json:"city"`
	Phone			string	`json:"phone"`
	AddressLine1	string	`json:"address_line_1"`
	AddressLine2	database.NullString	`json:"address_line_2,omitempty"`
	State			database.NullString	`json:"state,omitempty"`
	Country			string	`json:"country"`
	PostalCode		string	`json:"postal_code"`
	Territory		string	`json:"territory"`
	Employees		[]*Employee	`json:"employees",omitempty`
}

func (repository *Repository) GetOfficeByID(id uint64) (*Office, error) {
	row := repository.Conn.QueryRow("SELECT o.officeCode, o.city, o.phone," +
		"\no.addressLine1, o.addressLine2, o.state, o.country, o.postalCode, o.territory " +
		"\nFROM offices o WHERE o.officeCode=(?)", id)
	var officCode, city, phone, addressLine1, country, postalCode, territory string
	var addressLine2, state database.NullString
	switch err := row.Scan(&officCode, &city, &phone, &addressLine1, &addressLine2, &country,
		&state, &postalCode, &territory); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:

		office := &Office{
			OfficeCode:   officCode,
			City:         city,
			Phone:        phone,
			AddressLine1: addressLine1,
			AddressLine2: addressLine2,
			State:        state,
			Country:      country,
			PostalCode:   postalCode,
			Territory:    territory,
		}

		return office, nil
	default:
		return nil, err
	}
}

func (repository *Repository) GetOfficeEmployees(office *Office) ([]*Employee, error){
	var officeCode = office.OfficeCode

	rows, err := repository.Conn.Query("SELECT e.employeeNumber, e.lastName, e.firstName," +
		"\ne.extension, e.email, e.reportsTo, e.jobTitle FROM employees e" +
		"\nWHERE e.officeCode = (?)", officeCode)

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var employees []*Employee
	var employeeNumber, lastName, firstName, extension, email, jobTitle string
	var reportsTo database.NullString

	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&employeeNumber, &lastName, &firstName, &extension, &email, &reportsTo, &jobTitle)

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
		}

		employees = append(employees, employee)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return employees, nil
}
