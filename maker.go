package iagesdk

// Maker describe iAGE`s mobile vendor data structure
type Maker struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Models []string `json:"models"`
}
