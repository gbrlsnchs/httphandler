package httphandler

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) (interface{}, error) { return nil, nil }
	loggerFunc := func(err error, reqURI string) {}
	h := New(http.StatusOK, handlerFunc).SetLoggerFunc(loggerFunc)
	hClone := h.Clone()
	a := assert.New(t)
	h1, ok := h.(*handler)

	a.True(ok)

	h2, ok := hClone.(*handler)

	a.True(ok)
	a.NotEqual(h, hClone)
	a.NotEqual(h1, h2)
}

func TestNewHandler(t *testing.T) {
	a := assert.New(t)
	tests := []*struct {
		expected       interface{}
		expectedCode   int
		expectedErr    error
		expectedHeader http.Header
		h              Handler
	}{
		// #0
		{
			expected:       &struct{ Msg string }{Msg: "test"},
			expectedCode:   http.StatusOK,
			expectedErr:    nil,
			expectedHeader: globalHeader,
			h: New(http.StatusOK, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return &struct{ Msg string }{Msg: "test"}, nil
			}),
		},
		// #1
		{
			expected:       &struct{ Msg string }{Msg: "test"},
			expectedCode:   http.StatusAccepted,
			expectedErr:    NewError(http.StatusBadRequest, "test"),
			expectedHeader: globalHeader,
			h: New(http.StatusAccepted, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return &struct{ Msg string }{Msg: "test"}, NewError(http.StatusBadRequest, "test")
			}),
		},
		// #2
		{
			expected:       &struct{ Msg string }{Msg: "test"},
			expectedCode:   http.StatusCreated,
			expectedErr:    errors.New("test"),
			expectedHeader: globalHeader,
			h: New(http.StatusCreated, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return &struct{ Msg string }{Msg: "test"}, errors.New("test")
			}),
		},
		// #3
		{
			expected:       &struct{ Msg string }{Msg: "test"},
			expectedCode:   http.StatusOK,
			expectedErr:    nil,
			expectedHeader: http.Header{ContentType: []string{ContentTypeJSON}},
			h: New(http.StatusOK, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return &struct{ Msg string }{Msg: "test"}, nil
			}).JSON(),
		},
		// #4
		{
			expected:       &struct{ Msg string }{Msg: "test"},
			expectedCode:   http.StatusAccepted,
			expectedErr:    NewError(http.StatusBadRequest, "test"),
			expectedHeader: http.Header{ContentType: []string{ContentTypeXML}},
			h: New(http.StatusAccepted, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return &struct{ Msg string }{Msg: "test"}, NewError(http.StatusBadRequest, "test")
			}).XML(),
		},
	}

	for i, test := range tests {
		h, ok := test.h.(*handler)
		index := strconv.Itoa(i)

		a.True(ok, index)

		res, err := h.handlerFunc(nil, nil)

		a.Exactly(test.expected, res, index)
		a.Exactly(test.expectedCode, h.code, index)
		a.Exactly(test.expectedErr, err, index)
		a.Exactly(test.expectedHeader, h.header, index)
	}
}

func TestServeHTTP(t *testing.T) {
	type res struct {
		XMLName xml.Name `json:"-" xml:"test"`
		Message string   `json:"message" xml:"message"`
	}

	a := assert.New(t)
	tests := []*struct {
		contentType    string
		status         int
		response       *res
		err            error
		expected       interface{}
		expectedStatus int
	}{
		// #0
		{
			contentType:    ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #1
		{
			contentType:    ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #2
		{
			contentType:    ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(defaultErrCode, defaultErrMsg),
			expectedStatus: http.StatusBadRequest,
		},
		// #3
		{
			contentType:    ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #4
		{
			contentType:    ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #5
		{
			contentType:    ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(defaultErrCode, defaultErrMsg),
			expectedStatus: http.StatusBadRequest,
		},
		// #6
		{
			contentType:    ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #7
		{
			contentType:    ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #8
		{
			contentType:    ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(defaultErrCode, defaultErrMsg),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for i, test := range tests {
		h := New(test.status, func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			return test.response, test.err
		})

		if test.contentType == ContentTypeJSON {
			h.JSON()
		} else if test.contentType == ContentTypeXML {
			h.XML()
		}

		srv := httptest.NewServer(h)

		defer srv.Close()

		w, err := http.Get(srv.URL)
		index := strconv.Itoa(i)

		a.Nil(err, index)

		body, err := ioutil.ReadAll(w.Body)

		a.Nil(err, index)
		a.Exactly(test.expectedStatus, w.StatusCode, index)

		if test.contentType == ContentTypeTextPlain {
			a.Exactly(fmt.Sprintln(test.expected), string(body), index)

			continue
		}

		marshaller := json.Marshal

		if test.contentType == ContentTypeXML {
			marshaller = xml.Marshal
		}

		expected, err := marshaller(test.expected)

		a.Nil(err, index)
		a.Exactly(expected, body, index)
	}
}
