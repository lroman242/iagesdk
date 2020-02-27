package iagesdk

import "io"

// Request interface describe common API request to iAGE platform
type Request interface {
	URL() string
	Method() string
	Body() io.Reader
	//Validate() error
	//ToJson() string
}
