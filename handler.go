package httphandler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

// HandlerFunc is the accepted function to use in a Handler.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (Responder, error)

// LoggerFunc is the function that logs an error in a request.
type LoggerFunc func(err error, reqURI string)

// Handler is an HTTP handler that implements http.Handler.
type Handler struct {
	hfunc   HandlerFunc
	lfunc   LoggerFunc
	errCode int
	errMsg  string
	ctype   ContentType
}

// New creates a new Handler.
func New(hfunc HandlerFunc) *Handler {
	return &Handler{
		hfunc:   hfunc,
		errCode: http.StatusInternalServerError,
		errMsg:  http.StatusText(http.StatusInternalServerError),
		ctype:   ContentTypeTextPlain,
	}
}

// Clone deeply copies the caller Handler and returns the clone.
func (h *Handler) Clone() *Handler {
	return &Handler{
		hfunc: h.hfunc,
		lfunc: h.lfunc,
		ctype: h.ctype,
	}
}

// ServeHTTP runs the custom handler function and catches its
// error, if there is any.
//
// If the caught error implements Error, it is sent as a response
// serialized according to the "Content-Type" header.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.hfunc == nil {
		http.NotFound(w, r)

		return
	}

	res, err := h.hfunc(w, r)

	if err != nil {
		if h.lfunc != nil {
			go h.lfunc(err, r.RequestURI)
		}

		switch e := err.(type) {
		case Error:
			err = h.write(w, e, e.Status())

			if err != nil {
				http.Error(w, h.errMsg, h.errCode)
			}

		default:
			errWrapper := NewError(h.errCode, h.errMsg)
			err = h.write(w, errWrapper, errWrapper.Status())

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
		err = h.write(w, nil, res.Status())

	default:
		err = h.write(w, res.Body(), res.Status())
	}

	if err != nil {
		http.Error(w, h.errMsg, h.errCode)
	}
}

// SetErrCode sets the default Handler's error status code.
// If none is set, the default is http.StatusInternalServerError.
func (h *Handler) SetErrCode(c int) {
	h.errCode = c
}

// SetErrMsg sets the default Handler's error message.
// If none is set, the default is http.StatusText(http.StatusInternalServerError).
func (h *Handler) SetErrMsg(m string) {
	h.errMsg = m
}

// SetLoggerFunc sets a function for logging a caught error.
func (h *Handler) SetLoggerFunc(lfunc LoggerFunc) {
	h.lfunc = lfunc
}

// WithContentType returns the Handler instance
// with a "Content-Type" value already set.
func (h *Handler) WithContentType(ctype ContentType) *Handler {
	h.ctype = ctype

	return h
}

// write writes the response according to the status code and the
// "Content-Type" header set in the response's header.
func (h *Handler) write(w http.ResponseWriter, r interface{}, c int) error {
	w.Header().Set(ContentTypeHeader, string(h.ctype))
	w.WriteHeader(c)

	if r == nil {
		return nil
	}

	var b []byte
	var err error

	switch h.ctype {
	case ContentTypeJSON:
		b, err = json.Marshal(r)

	case ContentTypeMsgPack:
		fallthrough
	case ContentTypeXMsgPack:
		b, err = msgpack.Marshal(r)

	case ContentTypeXML:
		b, err = xml.Marshal(r)
	}

	if err != nil {
		return err
	}

	if h.ctype != ContentTypeTextPlain {
		_, err = w.Write(b)

		if err != nil {
			return err
		}

		return nil
	}

	_, err = fmt.Fprint(w, r)

	if err != nil {
		return err
	}

	return nil
}
