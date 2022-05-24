package models

type Attributes_values struct {
	Id         int    `json:"id_attributes"`
	Name       string `json:"name"`
	Attributes *Attributes
}
