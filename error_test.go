package httphandler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	a := assert.New(t)
	testErr := NewError(http.StatusBadRequest, "test")
	e, ok := testErr.(*httpErr)

	a.True(ok)
	a.Exactly(http.StatusBadRequest, e.Code)
	a.Exactly("test", e.Msg)
	a.Exactly(e.Code, testErr.Status())
	a.Exactly(e.Msg, testErr.Error())
}
