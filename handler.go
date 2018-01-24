package httphandler

import (
	"encoding/json"
	"net/http"
)

// HandlerFunc is the accepted function to use in a Handler.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (*Response, error)

// LoggerFunc is a function that logs an error in a request.
// Since its goal is only debugging errors, it runs in a different goroutine
// and passes a deep copy of the request as the first argument.
type LoggerFunc func(r http.Request, err error)

// ErrorHandlerFunc is a function used to handle runtime errors.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// MarshallerFunc is a function used to marshal data
// to be written by a response writer.
type MarshallerFunc func(v interface{}) ([]byte, error)

var (
	// DefaultContentType is the default Content-Type MIME
	// type a Handler uses when it is created.
	DefaultContentType = "application/json"
	// DefaultLoggerFunc is the default function used
	// by a Handler when it is created to log information
	// when an error occurs.
	DefaultLoggerFunc LoggerFunc
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
)

// Handler is an HTTP handler that implements http.Handler.
type Handler struct {
	ContentType      string
	HandlerFunc      HandlerFunc
	LoggerFunc       LoggerFunc
	ErrorHandlerFunc ErrorHandlerFunc
	MarshallerFunc   MarshallerFunc
	ErrMsg           string
	ErrCode          int
}

// New creates a new Handler with default settings.
func New(hfunc HandlerFunc) *Handler {
	return &Handler{
		ContentType:      DefaultContentType,
		HandlerFunc:      hfunc,
		LoggerFunc:       DefaultLoggerFunc,
		ErrorHandlerFunc: DefaultErrorHandlerFunc,
		MarshallerFunc:   DefaultMarshallerFunc,
		ErrCode:          DefaultErrCode,
		ErrMsg:           DefaultErrMsg,
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

	response, err := h.HandlerFunc(w, r)

	w.Header().Set("Content-Type", h.ContentType)

	if err != nil {
		switch e := err.(type) {
		case Error:
			err = h.write(w, e, e.Status())

			if err != nil {
				h.handleError(w, r, err)
			}
		default:
			h.handleError(w, r, err)
		}

		return
	}

	if response == nil {
		http.NotFound(w, r)

		return
	}

	err = h.write(w, response.Body, response.Code)

	if err != nil {
		h.handleError(w, r, err)
	}
}

func (h *Handler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	if h.LoggerFunc != nil {
		go h.LoggerFunc(*r, err)
	}

	if h.ErrorHandlerFunc != nil {
		h.ErrorHandlerFunc(w, r, err)
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
