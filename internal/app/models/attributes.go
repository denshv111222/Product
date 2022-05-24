package models

type Attributes struct {
	Id    int    `json:"id_attributes"`
	Name  string `json:"name"`
	Slug  string `json:"Slug"`
	Units *Units
}
