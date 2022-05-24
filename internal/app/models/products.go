package models

type Product struct {
	Id                int    `json:"Id"`
	Name              string `json:"Name"`
	Slug              string `json:"Slug"`
	Brand             *Brand
	SKU               string `json:"SKU"`
	Short_description string `json:"short_description"`
	Full_description  string `json:"full_description"`
	Sort              int    `json:"sort"`
}
