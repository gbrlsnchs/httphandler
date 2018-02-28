package httphandler_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gbrlsnchs/httphandler"
	. "github.com/gbrlsnchs/httphandler/internal"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		code           int
		marshallerFunc MarshallerFunc
		isErr          bool
	}{
		{
			code:           http.StatusOK,
			marshallerFunc: json.Marshal,
		},
		{
			code:           http.StatusBadRequest,
			marshallerFunc: json.Marshal,
			isErr:          true,
		},
		{
			code:           http.StatusOK,
			marshallerFunc: xml.Marshal,
		},
		{
			code:           http.StatusBadRequest,
			marshallerFunc: xml.Marshal,
			isErr:          true,
		},
		{
			code:           http.StatusNoContent,
			marshallerFunc: json.Marshal,
		},
		{
			code:           http.StatusNoContent,
			marshallerFunc: xml.Marshal,
		},
	}

	for i := range tt {
		test := tt[i]
		response := &DummyResponse{}

		if test.isErr {
			response.Data = DummyError
		} else {
			response.Data = DummyData
		}

		expected, err := test.marshallerFunc(response.Data)

		if want, got := (error)(nil), err; want != got {
			t.Errorf("want %v, got %v\n", want, got)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		h := New(func(w http.ResponseWriter, r *http.Request) (Responder, error) {
			if test.isErr {
				return nil, response
			}

			return response, nil
		})
		h.MarshallerFunc = test.marshallerFunc

		h.ServeHTTP(w, r)

		body := w.Body.Bytes()
		code := w.Code

		if want, got := expected, body; !bytes.Equal(want, got) {
			t.Errorf("want %s, got %s\n", string(want), string(got))
		}

		if want, got := response.Code, code; want != code {
			t.Errorf("want %d, got %d\n", want, got)
		}
	}
}
