package httphandler_test

import (
	"errors"
	"net/http"

	"github.com/gbrlsnchs/httphandler"
	"github.com/gbrlsnchs/httphandler/internal"
)

func Example() {
	h := httphandler.New(func(w http.ResponseWriter, _ *http.Request) (httphandler.Responder, error) {
		err := errors.New("Example error")

		if err != nil {
			return nil, &internal.DummyResponse{
				Data: internal.DummyError,
				Code: http.StatusBadRequest,
				Err:  err.Error(),
			}
		}

		return &internal.DummyResponse{
			Data: internal.DummyData,
			Code: http.StatusOK,
		}, nil
	})

	http.Handle("/example", h)
}
