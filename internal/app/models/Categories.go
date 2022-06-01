package models

type Categories struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Parent_id int    `json:"parent_id"`
}
