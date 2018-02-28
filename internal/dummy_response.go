package internal

type DummyResponse struct {
	Data interface{}
	Code int
	Err  string
}

func (d *DummyResponse) Body() interface{} {
	// return struct {
	// 	XMLName xml.Name `json:"-" xml:"error"`
	// 	Msg     string   `json:"message,omitempty" xml:"message"`
	// }{
	// 	Msg: r.Msg,
	// }
	return d.Data
}

func (d *DummyResponse) Error() string {
	return d.Err
}

func (d *DummyResponse) Status() int {
	return d.Code
}
