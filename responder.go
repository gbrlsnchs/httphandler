package httphandler

// Responder is an interface for wrapping
// an HTTP response with body and status code.
type Responder interface {
	// Body returns an HTTP message body.
	Body() interface{}
	// Status returns the respective
	// status code for an HTTP response.
	Status() int
}
