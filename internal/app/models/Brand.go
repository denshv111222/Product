package models

type Brand struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Slug string `json:"Slug"`
}

type BrandFilter struct {
	Brands []*Brand
	Filter *Filter
}
