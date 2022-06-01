package models

type Attributes struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"Slug"`
	Units *Units
}
