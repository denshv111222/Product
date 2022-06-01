package models

type Attributes_values struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Attributes *Attributes
}
