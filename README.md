# httphandler (Minimalist HTTP Handler)
[![Build Status](https://travis-ci.org/gbrlsnchs/httphandler.svg?branch=master)](https://travis-ci.org/gbrlsnchs/httphandler)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gbrlsnchs/httphandler)

## About
This package is a minimalist HTTP handler for [Go] (or Golang) HTTP servers.  
It is compatible with any response content type that has a marshaller function.

## Usage
Full documentation [here].

## Example
```go
package httphandler_test

import (
	"errors"
	"net/http"

	"github.com/gbrlsnchs/httphandler"
)

func Example() {
	h := httphandler.New(func(w http.ResponseWriter, _ *http.Request) (httphandler.Responder, error) {
		err := errors.New("Example error")

		if err != nil {
			return nil, &errorMockup{
				Msg:  err.Error(),
				Code: http.StatusBadRequest,
			}
		}

		return &responderMockup{
			msg:  "Hello, World!",
			code: http.StatusOK,
		}, nil
	})

	http.Handle("/example", h)
}
```

## Contribution
### How to help:
- Pull Requests
- Issues
- Opinions

[Go]: https://golang.org
[here]: https://godoc.org/github.com/gbrlsnchs/httphandler
