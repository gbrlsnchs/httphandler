package httphandler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	w := httptest.NewRecorder()

	Panic(w)

	a := assert.New(t)

	a.Exactly(defaultErrCode, w.Code)
	a.Exactly(defaultErrMsg, strings.TrimSuffix(w.Body.String(), "\n"))
}

func TestWrite(t *testing.T) {
	type res struct {
		XMLName xml.Name `json:"-" xml:"test"`
		Msg     string   `json:"message" xml:"message"`
	}

	a := assert.New(t)
	tests := []*struct {
		response            *res
		code                int
		contentType         string
		expected            interface{}
		expectedCode        int
		expectedContentType string
	}{
		// #0
		{
			response:            &res{Msg: "test"},
			code:                http.StatusOK,
			contentType:         ContentTypeTextPlain,
			expected:            &res{Msg: "test"},
			expectedCode:        http.StatusOK,
			expectedContentType: ContentTypeTextPlain,
		},
		// #1
		{
			response:            &res{Msg: "test"},
			code:                http.StatusOK,
			contentType:         ContentTypeJSON,
			expected:            &res{Msg: "test"},
			expectedCode:        http.StatusOK,
			expectedContentType: ContentTypeJSON,
		},
		// #2
		{
			response:            &res{Msg: "test"},
			code:                http.StatusOK,
			contentType:         ContentTypeXML,
			expected:            &res{Msg: "test"},
			expectedCode:        http.StatusOK,
			expectedContentType: ContentTypeXML,
		},
	}

	for i, test := range tests {
		w := httptest.NewRecorder()
		w.Header().Set(ContentType, test.contentType)
		index := strconv.Itoa(i)

		err := Write(w, test.code, test.response)

		a.Nil(err, index)
		a.Exactly(test.expectedCode, w.Code, index)
		a.Exactly(test.expectedContentType, w.Header().Get(ContentType), index)

		if test.contentType == ContentTypeTextPlain {
			a.Exactly(fmt.Sprintln(test.expected), w.Body.String(), index)

			continue
		}

		marshaller := json.Marshal

		if test.contentType == ContentTypeXML {
			marshaller = xml.Marshal
		}

		expected, err := marshaller(test.expected)

		a.Nil(err, index)
		a.Exactly(expected, w.Body.Bytes(), index)
	}
}
