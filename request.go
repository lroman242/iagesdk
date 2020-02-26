package iagesdk

import "io"

// V1Agencies is represent URI path for agencies API requests
const v1Agencies = `/v1/agencies/`

// Request interface describe common API request to iAGE platform
type Request interface {
	URL() string
	Method() string
	Body() io.Reader
	//Validate() error
	//ToJson() string
}
