package models

type Attributes_values_products struct {
	Produkt           *Product
	Attributes_values *Attributes_values
	Sort              int `json:"sort"`
}
