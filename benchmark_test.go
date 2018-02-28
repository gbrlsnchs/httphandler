package httphandler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gbrlsnchs/httphandler"
	. "github.com/gbrlsnchs/httphandler/internal"
)

func Benchmark(b *testing.B) {
	response := &DummyResponse{Data: DummyData, Code: http.StatusOK}
	h := New(func(w http.ResponseWriter, r *http.Request) (Responder, error) {
		return response, nil
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		h.ServeHTTP(w, r)
	}
}
