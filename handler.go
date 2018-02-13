package httphandler

import (
	"encoding/json"
	"net/http"
)

// HandlerFunc is the accepted function to use in a Handler.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (Responder, error)

// ErrorLoggerFunc is a function that logs an error in a request.
// Since its goal is only debugging errors, it runs in a different goroutine
// and passes a deep copy of the request as the first argument.
type ErrorLoggerFunc func(r http.Request, err error)

// ErrorHandlerFunc is a function used to handle runtime errors.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// MarshallerFunc is a function used to marshal data
// to be written by a response writer.
type MarshallerFunc func(v interface{}) ([]byte, error)

// RuntimeErrorFunc is a function used to return an
// Error interface which will be sent as response.
type RuntimeErrorFunc func(r http.Request) Error

var (
	// DefaultContentType is the default Content-Type MIME
	// type a Handler uses when it is created.
	DefaultContentType = "application/json"
	// DefaultErrorLoggerFunc is the default function used
	// by a Handler when it is created to log information
	// when an error occurs.
	DefaultErrorLoggerFunc ErrorLoggerFunc
	// DefaultErrorHandlerFunc is the default function for
	// handling runtime errors.
	DefaultErrorHandlerFunc ErrorHandlerFunc
	// DefaultMarshallerFunc is the default marshalling function
	// used by a Handler when it is created.
	DefaultMarshallerFunc = json.Marshal
	// DefaultErrMsg is the message used by a Handler
	// when it is created for setting a status code when an error occurs.
	DefaultErrMsg string
	// DefaultErrCode is the status used by a Handler
	// when it is created for setting a status code when an error occurs.
	DefaultErrCode int
	// DefaultRuntimeErrorFunc is the default function for
	// retrieving an Error interface that will be sent by
	// a Handler if a runtime error is caught.
	DefaultRuntimeErrorFunc RuntimeErrorFunc
)

// Handler is an HTTP handler that implements http.Handler.
type Handler struct {
	ContentType      string
	HandlerFunc      HandlerFunc
	ErrorLoggerFunc  ErrorLoggerFunc
	ErrorHandlerFunc ErrorHandlerFunc
	MarshallerFunc   MarshallerFunc
	ErrMsg           string
	ErrCode          int
	RuntimeErrorFunc RuntimeErrorFunc
}

// New creates a new Handler with default settings.
func New(hfunc HandlerFunc) *Handler {
	return &Handler{
		ContentType:      DefaultContentType,
		HandlerFunc:      hfunc,
		ErrorLoggerFunc:  DefaultErrorLoggerFunc,
		ErrorHandlerFunc: DefaultErrorHandlerFunc,
		MarshallerFunc:   DefaultMarshallerFunc,
		ErrCode:          DefaultErrCode,
		ErrMsg:           DefaultErrMsg,
		RuntimeErrorFunc: DefaultRuntimeErrorFunc,
	}
}

// ServeHTTP runs the custom handler function and catches its
// error, if there is any.
//
// If the caught error implements Error, it is sent as a response
// serialized using the marsheller function set.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.HandlerFunc == nil {
		http.NotFound(w, r)

		return
	}

	res, err := h.HandlerFunc(w, r)

	w.Header().Set("Content-Type", h.ContentType)

	if err != nil {
		h.logError(r, err)

		switch e := err.(type) {
		case Error:
			if err = h.write(w, e, e.Status()); err != nil {
				h.handleError(w, r, err)
			}
		default:
			h.handleError(w, r, err)
		}

		return
	}

	if res == nil {
		return
	}

	err = h.write(w, res.Body(), res.Status())

	if err != nil {
		h.logError(r, err)
		h.handleError(w, r, err)
	}
}

func (h *Handler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	if h.RuntimeErrorFunc != nil {
		runtimeErr := h.RuntimeErrorFunc(*r)

		if err = h.write(w, runtimeErr, runtimeErr.Status()); err != nil {
			h.logError(r, err)

			if h.ErrorHandlerFunc != nil {
				h.ErrorHandlerFunc(w, r, err)
			}
		}
	}
}

func (h *Handler) logError(r *http.Request, err error) {
	if h.ErrorLoggerFunc != nil {
		h.ErrorLoggerFunc(*r, err)
	}
}

func (h *Handler) write(w http.ResponseWriter, body interface{}, code int) error {
	var (
		b   []byte
		err error
	)

	if h.MarshallerFunc != nil {
		b, err = h.MarshallerFunc(body)

		if err != nil {
			return err
		}
	} else if s, ok := body.([]byte); ok {
		b = s
	} else if s, ok := body.(string); ok {
		b = []byte(s)
	}

	w.WriteHeader(code)

	if len(b) > 0 {
		if _, err = w.Write(b); err != nil {
			return err
		}
	}

	return nil
}
