package httphandler

import "net/http"

// HandlerFunc is an alias for the Handler's handling function.
type HandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, error)

// LoggerFunc is an alias for the the Handler's logging function.
type LoggerFunc func(err error, reqURI string)

// Handler implements http.Handler and catches errors returned in
// the handler function. The response is then sent as plain text, JSON or XML.
type Handler interface {
	http.Handler
	// Clone deep copies the caller Handler and returns the clone.
	Clone() Handler
	// Header provides the http.Header map contained inside the Handler.
	Header() http.Header
	// JSON sets the "Content-Type" header as "application/json".
	JSON() Handler
	// SetLoggerFunc registers a function for logging a caught error.
	SetLoggerFunc(LoggerFunc) Handler
	// XML sets the "Content-Type" header as "application/xml".
	XML() Handler
}

// handler represents an HTTP handler that implements both http.Handler and Handler.
type handler struct {
	code        int
	handlerFunc HandlerFunc
	header      http.Header
	loggerFunc  LoggerFunc
}

// New creates a new Handler.
//
// If all goes well, it writes a response according to the "Content-Type" header
// containing the returned value in its handling function.
func New(code int, handlerFunc HandlerFunc) Handler {
	h := make(http.Header)

	for k, v := range globalHeader {
		h[k] = v
	}

	return &handler{code: code, handlerFunc: handlerFunc, header: h}
}

// Clone deep copies the caller Handler and returns the clone.
func (h *handler) Clone() Handler {
	clone := &handler{
		code:        h.code,
		handlerFunc: h.handlerFunc,
		loggerFunc:  h.loggerFunc,
		header:      make(http.Header),
	}

	for k, v := range h.header {
		clone.header[k] = v
	}

	return clone
}

// Header provides the http.Header map contained inside the Handler.
func (h *handler) Header() http.Header {
	return h.header
}

// JSON sets the "Content-Type" header as "application/json".
func (h *handler) JSON() Handler {
	h.header.Set(ContentType, ContentTypeJSON)

	return h
}

// ServeHTTP runs the custom handling function and catches its
// error, if there is any.
//
// If the caught error implements Error, it will be sent as a response,
// serialized according to the "Content-Type" header.
//
// If Write returns an error, it calls Panic.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.handlerFunc == nil {
		return
	}

	for key := range h.header {
		w.Header().Set(key, h.header.Get(key))
	}

	res, err := h.handlerFunc(w, r)

	if err != nil {
		if h.loggerFunc != nil {
			go h.loggerFunc(err, r.RequestURI)
		}

		switch e := err.(type) {
		case Error:
			err = Write(w, e.Status(), e)

			if err != nil {
				Panic(w)
			}

		default:
			genericErr := NewError(defaultErrCode, defaultErrMsg)
			err = Write(w, genericErr.Status(), genericErr)

			if err != nil {
				Panic(w)
			}
		}

		return
	}

	err = Write(w, h.code, res)

	if err != nil {
		Panic(w)
	}
}

// SetLoggerFunc sets a function for logging a caught error.
func (h *handler) SetLoggerFunc(loggerFunc LoggerFunc) Handler {
	h.loggerFunc = loggerFunc

	return h
}

// XML sets the "Content-Type" header as "application/xml".
func (h *handler) XML() Handler {
	h.header.Set(ContentType, ContentTypeXML)

	return h
}
