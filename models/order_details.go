package models

type OrderDetails struct {
	Order				*Order		`json:"_"`
	Product				*Product	`json:"_"`
	QuantityOrdered		uint		`json:"quantity_ordered"`
	PriceEach			float64		`json:"price_each"`
	OrderLineNumber		string		`json:"order_line_number"`
}
