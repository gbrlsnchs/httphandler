package httphandler

// Error is the error captured by a Handler.
type Error interface {
	error
	Responder
}
