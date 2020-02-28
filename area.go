package iagesdk

// Area type describe iAGE`s geoArea sub type data structure for line items
type Area struct {
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Radius float64 `json:"radius"`
}
