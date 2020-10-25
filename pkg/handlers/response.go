package handlers

// ResponseMeta a struct of non resource related information that can
// help integrators to determine success or get some more technical message
// regarding errors
type ResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// ResponseErrors a slice of strings that represents errors which might be related to the
// request
type ResponseErrors []string

// Response is the basic structure of a response from the service
type Response struct {
	Meta   ResponseMeta   `json:"meta"`
	Errors ResponseErrors `json:"errors,omitempty"`
	Data   interface{}    `json:"data,omitempty"`
}

// AddError injects the stringified error into the response errors
// slice
func (dr *Response) AddError(err error) {
	dr.Errors = append(dr.Errors, err.Error())
}
