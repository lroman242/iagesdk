package iagesdk

// ErrorResponse describe general response error that follows a similar layout to iAGE's response
// errors, i.e. either a single message
type ErrorResponse struct {
	Error string `json:"error"`
}
