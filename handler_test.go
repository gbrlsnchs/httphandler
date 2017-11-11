package httphandler_test

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

	. "github.com/gbrlsnchs/httphandler"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
)

func TestClone(t *testing.T) {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) (Responder, error) { return nil, nil }
	loggerFunc := func(err error, reqURI string) {}
	h := New(handlerFunc)

	h.SetLoggerFunc(loggerFunc)

	hClone := h.Clone()
	a := assert.New(t)

	a.NotEqual(h, hClone)
}

func TestNew(t *testing.T) {
	type res struct {
		XMLName xml.Name `json:"-" msgpack:"-" xml:"test"`
		Message string   `json:"message" msgpack:"message" xml:"message"`
	}

	a := assert.New(t)
	tests := []*struct {
		ctype          ContentType
		status         int
		response       *res
		err            error
		expected       interface{}
		expectedStatus int
	}{
		// #0
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #1
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #2
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			expectedStatus: http.StatusInternalServerError,
		},
		// #3
		{
			ctype:          ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #4
		{
			ctype:          ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #5
		{
			ctype:          ContentTypeJSON,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			expectedStatus: http.StatusInternalServerError,
		},
		// #6
		{
			ctype:          ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #7
		{
			ctype:          ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #8
		{
			ctype:          ContentTypeXML,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			expectedStatus: http.StatusInternalServerError,
		},
		// #9
		{
			ctype:          ContentTypeMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #10
		{
			ctype:          ContentTypeMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #11
		{
			ctype:          ContentTypeMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			expectedStatus: http.StatusInternalServerError,
		},
		// #12
		{
			ctype:          ContentTypeXMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusOK,
		},
		// #13
		{
			ctype:          ContentTypeXMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            NewError(http.StatusBadRequest, "test"),
			expected:       NewError(http.StatusBadRequest, "test"),
			expectedStatus: http.StatusBadRequest,
		},
		// #14
		{
			ctype:          ContentTypeXMsgPack,
			status:         http.StatusOK,
			response:       &res{Message: "test"},
			err:            errors.New("test"),
			expected:       NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			expectedStatus: http.StatusInternalServerError,
		},
		// #15
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusSwitchingProtocols,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       nil,
			expectedStatus: http.StatusSwitchingProtocols,
		},
		// #16
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusProcessing,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       nil,
			expectedStatus: http.StatusProcessing,
		},
		// #17
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusCreated,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusCreated,
		},
		// #18
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusNoContent,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       nil,
			expectedStatus: http.StatusNoContent,
		},
		// #19
		{
			ctype:          ContentTypeTextPlain,
			status:         http.StatusResetContent,
			response:       &res{Message: "test"},
			err:            nil,
			expected:       &res{Message: "test"},
			expectedStatus: http.StatusResetContent,
		},
	}

	for i, test := range tests {
		h := New(func(w http.ResponseWriter, r *http.Request) (Responder, error) {
			return NewResponder(test.response, test.status), test.err
		}).WithContentType(test.ctype)
		srv := httptest.NewServer(h)

		defer srv.Close()

		w, err := http.Get(srv.URL)
		index := strconv.Itoa(i)

		a.Nil(err, index)
		a.IsType(ContentType(""), test.ctype, index)

		body, err := ioutil.ReadAll(w.Body)

		t.Logf("#%s body = %s\n", index, body)

		a.Nil(err, index)
		a.Exactly(test.expectedStatus, w.StatusCode, index)

		if test.ctype == ContentTypeTextPlain {
			if test.expected == nil {
				a.Exactly("", string(body), index)

				continue
			}

			a.Exactly(fmt.Sprint(test.expected), string(body), index)

			continue
		}

		expected, err := func() ([]byte, error) {
			if test.ctype == ContentTypeJSON {
				return json.Marshal(test.expected)
			}

			if test.ctype == ContentTypeXML {
				return xml.Marshal(test.expected)
			}

			return msgpack.Marshal(test.expected)
		}()

		a.Nil(err, index)

		if test.expected != nil {
			a.Exactly(expected, body, index)
		}
	}
}
