package httphandler

// Response is the response a Handler sends.
type Response struct {
	Body interface{}
	Code int
}
