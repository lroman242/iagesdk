package iagesdk

// Category describe iAGE`s category data structure
type Category struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NameRu string `json:"nameRu"`
	Code   string `json:"code"`
}
