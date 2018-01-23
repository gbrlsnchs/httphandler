package httphandler

import "encoding/xml"

// Error is the error returned by a Handler.
type Error interface {
	error
	// Status returns the Error's HTTP status.
	Status() int
}

// httpErr is a helper that creates an Error
// with JSON, MsgPack and XML support.
type httpErr struct {
	XMLName xml.Name `json:"-" msgpack:"-" xml:"error"`
	Code    int      `json:"code,omitempty" msgpack:"code,omitempty" xml:"code,omitempty"`
	Msg     string   `json:"message,omitempty" msgpack:"message,omitempty" xml:"message,omitempty"`
}

// NewError creates an Error with a status code and a friendly message.
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
