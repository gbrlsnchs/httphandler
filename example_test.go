package httphandler_test

import (
	"net/http"

	"github.com/gbrlsnchs/httphandler"
)

func Example() {
	h := httphandler.New(func(w http.ResponseWriter, _ *http.Request) (*httphandler.Response, error) {
		return &httphandler.Response{
			Body: struct {
				Msg string `json:"message"`
			}{"Hello, World!"},
			Code: http.StatusOK,
		}, nil
	})

	http.Handle("/example", h)
}
