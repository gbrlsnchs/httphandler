package httphandler

// Error is the error captured by a Handler.
type Error interface {
	error
	Status() int // Status returns the Error's HTTP status.
}
