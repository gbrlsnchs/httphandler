package httphandler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func write(w http.ResponseWriter, resp interface{}, statusCode int) error {
	header := w.Header()
	ctype := header.Get(ContentType)

	w.WriteHeader(statusCode)

	if resp == nil {
		return nil
	}

	var err error
	var b []byte

	switch ctype {
	case ContentTypeJSON:
		b, err = json.Marshal(resp)

	case ContentTypeMsgPack:
		fallthrough
	case ContentTypeXMsgPack:
		b, err = msgpack.Marshal(resp)

	case ContentTypeXML:
		b, err = xml.Marshal(resp)

	default:
		ctype = ContentTypeTextPlain

	}

	if err != nil {
		return err
	}

	if ctype != ContentTypeTextPlain {
		_, err = w.Write(b)

		if err != nil {
			return err
		}

		return nil
	}

	_, err = fmt.Fprint(w, resp)

	if err != nil {
		return err
	}

	return nil
}
