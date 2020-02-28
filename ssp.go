package iagesdk

// SSP describe iAGE`s SSP data structure
type SSP struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	UID                  string `json:"uid"`
	Disabled             bool   `json:"disabled"`
	ClickMacrosUnescaped string `json:"clickMacrosUnescaped"`
	ClickMacrosEscaped   string `json:"clickMacrosEscaped"`
}
