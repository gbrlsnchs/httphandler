package httphandler_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gbrlsnchs/httphandler"
)

func TestHandlerJSONResponse(t *testing.T) {
	response := &Response{Body: &responseMockup{Msg: "test"}, Code: http.StatusOK}
	expectedResponse, err := json.Marshal(response.Body)
	expectedCode := response.Code

	if err != nil {
		t.Errorf("%v\n", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return response, nil
	})

	h.ServeHTTP(w, r)

	body := w.Body.Bytes()
	code := w.Code

	if !bytes.Equal(expectedResponse, body) {
		t.Errorf("%s is not expected response (%s)\n", string(expectedResponse), string(body))
	}

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}

func TestHandlerJSONResponseWithError(t *testing.T) {
	responseErr := &errorMockup{Msg: "Oops!", Code: http.StatusBadRequest}
	expectedResponse, err := json.Marshal(responseErr)
	expectedCode := responseErr.Code

	if err != nil {
		t.Errorf("%v\n", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return nil, responseErr
	})

	h.ServeHTTP(w, r)

	body := w.Body.Bytes()
	code := w.Code

	if !bytes.Equal(expectedResponse, body) {
		t.Errorf("%s is not expected response (%s)\n", string(expectedResponse), string(body))
	}

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}

func TestHandlerXMLResponse(t *testing.T) {
	response := &Response{Body: &responseMockup{Msg: "test"}, Code: http.StatusOK}
	expectedResponse, err := xml.Marshal(response.Body)
	expectedCode := response.Code

	if err != nil {
		t.Errorf("%v\n", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return response, nil
	})
	h.MarshallerFunc = xml.Marshal

	h.ServeHTTP(w, r)

	body := w.Body.Bytes()
	code := w.Code

	if !bytes.Equal(expectedResponse, body) {
		t.Errorf("%s is not expected response (%s)\n", string(expectedResponse), string(body))
	}

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}

func TestHandlerXMLResponseWithError(t *testing.T) {
	responseErr := &errorMockup{Msg: "Oops!", Code: http.StatusBadRequest}
	expectedResponse, err := xml.Marshal(responseErr)
	expectedCode := responseErr.Code

	if err != nil {
		t.Errorf("%v\n", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return nil, responseErr
	})
	h.MarshallerFunc = xml.Marshal

	h.ServeHTTP(w, r)

	body := w.Body.Bytes()
	code := w.Code

	if !bytes.Equal(expectedResponse, body) {
		t.Errorf("%s is not expected response (%s)\n", string(expectedResponse), string(body))
	}

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}

func TestHandlerWithoutResponse(t *testing.T) {
	expectedCode := http.StatusNotFound
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return nil, nil
	})

	h.ServeHTTP(w, r)

	code := w.Code

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}

func TestHandlerNoContent(t *testing.T) {
	expectedCode := http.StatusNoContent
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return &Response{Code: http.StatusNoContent}, nil
	})

	h.ServeHTTP(w, r)

	code := w.Code

	if expectedCode != code {
		t.Errorf("%d is not expected status (%d)\n", expectedCode, code)
	}
}
