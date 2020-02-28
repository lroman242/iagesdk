package iagesdk

// OS describe iAGE`s operating system data structure
type OS struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}
