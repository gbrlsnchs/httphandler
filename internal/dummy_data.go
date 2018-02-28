package internal

import "encoding/xml"

var DummyData = struct {
	XMLName xml.Name `json:"-" xml:"response"`
	Msg     string   `json:"message" xml:"message"`
}{Msg: "Hello, world!"}
