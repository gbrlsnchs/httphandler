package httphandler_test

import (
	"errors"
	"net/http"

	"github.com/gbrlsnchs/httphandler"
)

func Example() {
	h := httphandler.New(func(w http.ResponseWriter, _ *http.Request) (httphandler.Responder, error) {
		err := errors.New("Example error")

		if err != nil {
			return nil, &errorMockup{
				Msg:  err.Error(),
				Code: http.StatusBadRequest,
			}
		}

		return &responderMockup{
			msg:  "Hello, World!",
			code: http.StatusOK,
		}, nil
	})

	http.Handle("/example", h)
}
