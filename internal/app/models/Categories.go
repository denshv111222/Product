package models

type Categories struct {
	Id        int    `json:"id_categories"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Parent_id string `json:"parent_id"`
}
