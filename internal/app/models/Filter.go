package models

type FieldFilter struct {
	Field      string `json:"field"`
	Operations string `json:"operations""`
	Value      string `json:"value"`
}

type Filter struct {
	Fields *[]FieldFilter
	Sorts  *[]FieldSort
	Pages  *Pages
}
type FieldSort struct {
	Sort     string `json:"sort"`
	SortView string `json:"sortview"`
}

type Pages struct {
	СurrentPage        int `json:"currentpage"`
	AllPages           int `json:"allpages""`
	CountsRecordOnPage int `json:"countsrecordonpage"`
	AllRecords         int `json:"allrecords"`
	RemainedRecords    int `json:"remainedrecords"`
}
