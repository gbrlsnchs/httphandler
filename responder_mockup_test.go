package httphandler_test

import "encoding/xml"

type responderMockup struct {
	code int
	msg  string
}

func (r *responderMockup) Body() interface{} {
	return struct {
		XMLName xml.Name `json:"-" xml:"error"`
		Msg     string   `json:"message,omitempty" xml:"message"`
	}{
		Msg: r.msg,
	}
}

func (r *responderMockup) Status() int {
	return r.code
}
