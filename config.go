package httphandler

import "net/http"

const (
	ContentType          = "Content-Type"
	ContentTypeJSON      = "application/json"
	ContentTypeTextPlain = "text/plain"
	ContentTypeXML       = "application/xml"
)

var (
	// defaultErrCode is the default code for when everything fails
	defaultErrCode = http.StatusInternalServerError
	// defaultErrMsg is the default message for when everything fails
	defaultErrMsg = http.StatusText(http.StatusInternalServerError)
	// globalHeader is an http.Header map that is copied to every Handler's
	// own http.Header map.
	globalHeader = make(http.Header)
)

func init() {
	// This is done because Set does more than only setting the value.
	globalHeader.Set(ContentType, ContentTypeTextPlain)
}

// Config sets the default error code and the default error message to
// throw when Panic is called or when an error that doesn't implement Error
// is caught by a Handler.
//
// If the second parameter is an empty string,
// the default error message is the default error code
// parsed by the http.StatusText function.
func Config(code int, msg string) {
	defaultErrCode = code

	if msg == "" {
		defaultErrMsg = http.StatusText(code)

		return
	}

	defaultErrMsg = msg
}

// Header returns the global header map.
func Header() http.Header {
	return globalHeader
}
