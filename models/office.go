package models


type Office struct {
	OfficeCode		string	`json:"office_code"`
	City			string	`json:"city"`
	Phone			uint	`json:"phone"`
	AddressLine1	string	`json:"address_line_1"`
	AddressLine2	string	`json:"address_line_2,omitempty"`
	State			string	`json:"state,omitempty"`
	Country			string	`json:"country"`
	PostalCode		string	`json:"postal_code"`
	Territory		string	`json:"territory"`
}
