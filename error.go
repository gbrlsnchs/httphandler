package httphandler

import "encoding/xml"

// Error is a custom error that is sent as a response by the Write function.
type Error interface {
	error
	// Status prints the Error's status code.
	Status() int
}

// httpErr is the default struct used by this package for
// error responses.
type httpErr struct {
	XMLName xml.Name `json:"-" xml:"error"`
	Code    int      `json:"code,omitempty" xml:"code,omitempty"`
	Msg     string   `json:"message,omitempty xml:"message,omitempty`
}

// NewError creates a new custom error that gets caught
// by the Handler if returned by its handling function.
func NewError(code int, msg string) Error {
	return &httpErr{Code: code, Msg: msg}
}

// Error prints the Error's message.
func (e *httpErr) Error() string {
	return e.Msg
}

// Status prints the Error's status code.
func (e *httpErr) Status() int {
	return e.Code
}
