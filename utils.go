package httphandler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Panic short-circuits the Handler and sends a generic error to the client.
//
// It is a utility function that can be used outside a Handler.
func Panic(w http.ResponseWriter) {
	http.Error(w, defaultErrMsg, defaultErrCode)
}

// Write writes the response. The default response is sent as plain text,
// but one is able to set the "Content-Type" header accordingly to send a JSON or XML response.
//
// It is a utility function that can be used outside a Handler.
func Write(w http.ResponseWriter, statusCode int, res interface{}) error {
	header := w.Header()
	cType := header.Get(ContentType)

	var err error
	var r []byte

	switch cType {
	case ContentTypeJSON:
		r, err = json.Marshal(res)

	case ContentTypeXML:
		r, err = xml.Marshal(res)

	default:
		cType = ContentTypeTextPlain

	}

	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)

	if cType != ContentTypeTextPlain {
		_, err = w.Write(r)

		if err != nil {
			return err
		}

		return nil
	}

	fmt.Fprintln(w, res)

	return nil
}
