package models

type ProductLine struct {
	ProductLine			string		`json:"product_line"`
	TextDescription		string		`json:"text_description,omitempty"`
	HtmlDescription		string		`json:"html_description,omitempty"`
	Image				[]byte		`json:"image,omitempty"`
}
