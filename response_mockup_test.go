package httphandler_test

import "encoding/xml"

type responseMockup struct {
	XMLName xml.Name `json:"-" xml:"response"`
	Msg     string   `json:"message,omitempty" xml:"message"`
}
