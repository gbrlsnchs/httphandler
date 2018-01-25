package httphandler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gbrlsnchs/httphandler"
)

func Benchmark(b *testing.B) {
	b.ReportAllocs()

	h := New(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return &Response{Body: &struct{ test string }{test: "foobar"}, Code: http.StatusOK}, nil
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(w, r)
	}
}
