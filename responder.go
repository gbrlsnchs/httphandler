package httphandler

// Responder is the value returned by a Handler.
type Responder interface {
	// Content returns an interface to be read by a Handler.
	Content() interface{}
	// Status returns the status read by a Handler.
	Status() int
}

// responder is a helping wrapper.
type responder struct {
	content interface{}
	code    int
}

// NewResponder returns a wrapper to be used in a Handler.
func NewResponder(content interface{}, code int) Responder {
	return &responder{content: content, code: code}
}

// Content returns the wrapper's content.
func (r *responder) Content() interface{} {
	return r.content
}

// Status returns the wrapper's status.
func (r *responder) Status() int {
	return r.code
}
