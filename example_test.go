package httphandler_test

import (
	"encoding/xml"
	"net/http"

	"github.com/gbrlsnchs/httphandler"
)

func Example(err error) {
	type myResponse struct {
		XMLName xml.Name `json:"-" msgpack:"-" xml:"myResponse"`
		Msg     string   `json:"message" msgpack:"message" xml:"message"`
	}

	h := httphandler.New(func(w http.ResponseWriter, r *http.Request) (httphandler.Responder, error) {
		// ...
		if err != nil {
			// httphandler.NewError is used here,
			// but anything that implements the
			// httphandler.Error interface will be sent as a response
			return nil, httphandler.NewError(http.StatusBadRequest, "Oops!")
		}

		res := &myResponse{
			Msg: "foobar",
		}

		// The httphandler.NewResponder wrapper can be used as much as
		// a struct/interface that implements the Responder interface as a return value.
		return httphandler.NewResponder(res, http.StatusOK), nil
	})

	j := h.Clone()
	m := h.Clone()
	x := h.Clone()

	http.Handle("/plain", h)
	http.Handle("/json", j.SetJSON())
	http.Handle("/msg-pack", m.SetMsgPack())
	http.Handle("/xml", x.SetXML())
}
