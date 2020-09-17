package models

import (
	"database/sql"
	"fmt"
	"gobdd/database"
	"time"
)

// Category struct
type Customer struct {
	CustomerNumber			string	`json:"customer_number"`
	CustomerName        	string	`json:"customer_name"`
	ContactLastName			string	`json:"contact_last_name"`
	ContactFirstName   		string	`json:"contact_first_name"`
	Phone					string	`json:"phone"`
	AddressLine1			string	`json:"address_line_1"`
	AddressLine2			database.NullString	`json:"address_line_2,omitempty"`
	City					string	`json:"city"`
	State					database.NullString	`json:"state,omitempty"`
	PostalCode				database.NullString 	`json:"postal_code"`
	Country 				string 	`json:"country"`
	SalesRepEmployee	 	*Employee `json:"_,omitempty"`
	CreditLimit 			database.NullString 	`json:"credit_limit,omitempty"`
	Orders					[]*Order	`json:"orders,omitempty"`

}

func (repository *Repository) GetCustomerByID(id uint64) (*Customer, error) {
	row := repository.Conn.QueryRow("SELECT c.customerNumber, c.customerName, c.contactLastName," +
		"c.contactFirstName, c.phone, c.addressLine1, c.addressLine2, c.city, c.state, c.postalCode, c.country, " +
		"c.creditLimit FROM customers c WHERE c.customerNumber=(?)", id)
	var customerNumber, customerName, contactLastName, contactFirstName, phone, addressLine1, city, country string
	var addressLine2, state, creditLimit, postalCode database.NullString
	switch err := row.Scan(&customerNumber, &customerName, &contactLastName, &contactFirstName, &phone, &addressLine1,
		&addressLine2, &city, &state, &postalCode, &country, &creditLimit); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:

		customer := &Customer{
			CustomerNumber:   customerNumber,
			CustomerName:     customerName,
			ContactLastName:  contactLastName,
			ContactFirstName: contactFirstName,
			Phone:            phone,
			AddressLine1:     addressLine1,
			AddressLine2:     addressLine2,
			City:             city,
			State:            state,
			PostalCode: 	  postalCode,
			Country:          country,
			SalesRepEmployee: nil,
			CreditLimit:      creditLimit,
		}
		return customer, nil
	default:
		return nil, err
	}
}

func (repository *Repository) GetCustomerOrders(customer *Customer) ([]*Order, error){
	var customerNumber = customer.CustomerNumber

	rows, err := repository.Conn.Query("SELECT o.orderNumber, o.orderDate, o.requiredDate," +
		"\no.shippedDate, o.status, o.comments, SUM(od.priceEach * od.quantityOrdered) as total FROM orders o" +
		"\nINNER JOIN orderdetails od on od.orderNumber = o.orderNumber" +
		"\nWHERE o.customerNumber = (?)" +
		"\nGROUP BY o.orderNumber", customerNumber)

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var orders []*Order
	var total float64
	var orderNumber, status string
	var orderDate, requiredDate time.Time
	var comments database.NullString
	var shippedDate database.NullTime

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&orderNumber, &orderDate, &requiredDate, &shippedDate, &status, &comments, &total)

		if err != nil {
			return nil, fmt.Errorf("could not get articles : %v", err)
		}
		order := &Order{
			OrderNumber: orderNumber,
			OrderDate:    orderDate,
			RequiredDate: requiredDate,
			ShippedDate:  shippedDate,
			Status:       status,
			Comments:     comments,
			TotalPrice:   total,
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return orders, nil
}