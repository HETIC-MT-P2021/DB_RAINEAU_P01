package models

import (
	"fmt"
	"gobdd/database"
	"time"
)

type Order struct {
	OrderNumber			string		`json:"order_number"`
	OrderDate			time.Time	`json:"order_date"`
	RequiredDate		time.Time	`json:"required_date"`
	ShippedDate			database.NullTime	`json:"shipped_date,omitempty"`
	Status				string		`json:"status"`
	Comments			database.NullString	`json:"comments,omitempty"`
	TotalPrice			float64		`json:"total,omitempty"`
}

func (repository *Repository) GetOrderProductsById(id uint64) ([]*Product, error) {
	rows, err := repository.Conn.Query("SELECT p.productCode, p.productName," +
		"\np.productScale, p.productVendor, p.productDescription, p.quantityInStock, p.buyPrice, p.MSRP " +
		"\nFROM products p" +
		"\nINNER JOIN orderdetails od ON od.productCode = p.productCode" +
		"\nWHERE od.orderNumber =(?)", id)

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var products []*Product
	var productCode, productName, productScale, productVendor, productDescription string
	var quantityInStock uint
	var buyPrice, msrp float64

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&productCode, &productName, &productScale, &productVendor, &productDescription, &quantityInStock,
			&buyPrice, &msrp)

		if err != nil {
			return nil, fmt.Errorf("could not get articles : %v", err)
		}
		product := &Product{
			ProductCode:        productCode,
			ProductName:        productName,
			ProductScale:       productScale,
			ProductVendor:      productVendor,
			ProductDescription: productDescription,
			QuantityInStock:    quantityInStock,
			BuyPrice:           buyPrice,
			Msrp:               msrp,
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return products, nil
}
