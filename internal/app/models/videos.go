package models

type Videos struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Storage string `json:"storage"`
	Path    string `json:"path"`
}
