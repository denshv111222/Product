package models

type Images struct {
	Id      int    `json:"id_images"`
	Name    string `json:"name"`
	Storage string `json:"storage"`
	Path    string `json:"path"`
}
