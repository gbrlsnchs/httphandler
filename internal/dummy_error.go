package internal

import "encoding/xml"

var DummyError = struct {
	XMLName xml.Name `json:"-" xml:"error"`
	Msg     string   `json:"message" xml:"message"`
}{Msg: "Oops!"}
