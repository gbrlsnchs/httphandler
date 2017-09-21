package httphandler_test

import (
	"encoding/xml"
	"net/http"

	"github.com/gbrlsnchs/httphandler"
)

func Example(err error) {
	type myResponse struct {
		XMLName xml.Name `json:"-" xml:"myResponse"`
		Msg     string   `json:"message" xml:"message"`
	}

	h := httphandler.New(http.StatusOK, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		// ...
		if err != nil {
			// httphandler.NewError is used here,
			// but anything that implements the
			// httphandler.Error interface will be sent as a response
			return nil, httphandler.NewError(http.StatusBadRequest, "Oops!")
		}

		return &myResponse{Msg: "Hello World"}, nil
	})

	c := h.Clone()

	http.Handle("/json", h.JSON())
	http.Handle("/xml", c.XML())
}

func ExampleWrite(w http.ResponseWriter, myResponse interface{}, err error) {
	if err != nil {
		e := httphandler.NewError(http.StatusBadRequest, err.Error())
		err = httphandler.Write(w, e.Status(), e)

		if err != nil {
			httphandler.Panic(w)
		}

		return
	}

	err = httphandler.Write(w, http.StatusOK, myResponse)

	if err != nil {
		httphandler.Panic(w)
	}
}
