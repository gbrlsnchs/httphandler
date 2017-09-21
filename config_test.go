package httphandler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	Config(http.StatusBadRequest, "test")

	a := assert.New(t)

	a.Exactly(http.StatusBadRequest, defaultErrCode)
	a.Exactly("test", defaultErrMsg)
}

func TestHeader(t *testing.T) {
	assert.Exactly(t, globalHeader, Header())
}
