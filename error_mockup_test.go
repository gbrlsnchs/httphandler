package httphandler_test

import "encoding/xml"

type errorMockup struct {
	XMLName xml.Name `json:"-" xml:"error"`
	Msg     string   `json:"message" xml:"message"`
	Code    int      `json:"code" xml:"code"`
}

func (e *errorMockup) Body() interface{} {
	return e
}

func (e *errorMockup) Error() string {
	return e.Msg
}

func (e *errorMockup) Status() int {
	return e.Code
}
