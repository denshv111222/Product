package models

type Categories_products struct {
	Categories *Categories
	Product    *Product
	Sort       int `json:"sort"`
}
