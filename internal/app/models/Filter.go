package models

type Field struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
	Order     bool   `json:"order"` // true for ASC false for DESC
}

type PageRequest struct {
	PageNumber   int `json:"pg_number"`
	PageLength   int `json:"pg_length"`
	Fields       *[]Field
	TotalRecords int `json:"total_rec"`
}
