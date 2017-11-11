package httphandler

// Responder is the value returned by a Handler.
type Responder interface {
	// Body returns an interface to be read by a Handler.
	Body() interface{}
	// Status returns the status read by a Handler.
	Status() int
}

// responder is a helping wrapper.
type responder struct {
	body interface{}
	code int
}

// NewResponder returns a wrapper to be used in a Handler.
func NewResponder(b interface{}, c int) Responder {
	return &responder{body: b, code: c}
}

// Content returns the wrapper's content.
func (r *responder) Body() interface{} {
	return r.body
}

// Status returns the wrapper's status.
func (r *responder) Status() int {
	return r.code
}
