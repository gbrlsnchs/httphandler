package httphandler

import "net/http"

// HandlerFunc is the accepted function to use in a Handler.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (Responder, error)

// LoggerFunc is the function that logs an error in a request.
type LoggerFunc func(err error, reqURI string)

// Handler is an HTTP handler that implements http.Handler.
type Handler struct {
	handlerFunc HandlerFunc
	header      http.Header
	loggerFunc  LoggerFunc
	errCode     int
	errMsg      string
}

// New creates a new Handler.
func New(handlerFunc HandlerFunc) *Handler {
	h := make(http.Header)

	h.Set(ContentType, ContentTypeTextPlain)

	internalErr := http.StatusInternalServerError

	return &Handler{
		handlerFunc: handlerFunc,
		header:      h,
		errCode:     internalErr,
		errMsg:      http.StatusText(internalErr),
	}
}

// Clone deeply copies the caller Handler and returns the clone.
func (h *Handler) Clone() *Handler {
	clone := &Handler{
		handlerFunc: h.handlerFunc,
		loggerFunc:  h.loggerFunc,
		header:      make(http.Header),
	}

	for k, v := range h.header {
		clone.header[k] = v
	}

	return clone
}

// Header returns the Handler's http.Header.
func (h *Handler) Header() http.Header {
	return h.header
}

// ServeHTTP runs the custom handler function and catches its
// error, if there is any.
//
// If the caught error implements Error, it is sent as a response
// serialized according to the "Content-Type" header.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.handlerFunc == nil {
		http.NotFound(w, r)

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
			err = write(w, e, e.Status())

			if err != nil {
				http.Error(w, h.errMsg, h.errCode)
			}

		default:
			errWrapper := NewError(h.errCode, h.errMsg)
			err = write(w, errWrapper, errWrapper.Status())

			if err != nil {
				http.Error(w, h.errMsg, h.errCode)
			}
		}

		return
	}

	switch res.Status() {
	case http.StatusContinue:
		fallthrough

	case http.StatusSwitchingProtocols:
		fallthrough

	case http.StatusProcessing:
		fallthrough

	case http.StatusNoContent:
		err = write(w, nil, res.Status())

	default:
		err = write(w, res.Content(), res.Status())
	}

	if err != nil {
		http.Error(w, h.errMsg, h.errCode)
	}
}

// SetErrCode sets the default Handler's error status code.
// If none is set, the default is http.StatusInternalServerError.
func (h *Handler) SetErrCode(c int) *Handler {
	h.errCode = c

	return h
}

// SetErrMsg sets the default Handler's error message.
// If none is set, the default is http.StatusText(http.StatusInternalServerError).
func (h *Handler) SetErrMsg(m string) *Handler {
	h.errMsg = m

	return h
}

// SetJSON sets the "Content-Type" header as "application/json".
func (h *Handler) SetJSON() *Handler {
	h.header.Set(ContentType, ContentTypeJSON)

	return h
}

// SetLoggerFunc sets a function for logging a caught error.
func (h *Handler) SetLoggerFunc(loggerFunc LoggerFunc) *Handler {
	h.loggerFunc = loggerFunc

	return h
}

// SetMsgPack sets the "Content-Type" header as "application/vnd.msgpack".
func (h *Handler) SetMsgPack() *Handler {
	h.header.Set(ContentType, ContentTypeMsgPack)

	return h
}

// SetXMsgPack sets the "Content-Type" header as "application/x-msgpack".
func (h *Handler) SetXMsgPack() *Handler {
	h.header.Set(ContentType, ContentTypeXMsgPack)

	return h
}

// SetXML sets the "Content-Type" header as "application/xml".
func (h *Handler) SetXML() *Handler {
	h.header.Set(ContentType, ContentTypeXML)

	return h
}
