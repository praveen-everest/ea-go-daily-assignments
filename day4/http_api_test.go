package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hello", nil)

	helloHandler(recorder, r)

	result := recorder.Result().Body
	data, _ := io.ReadAll(result)

	assert.Equal(t, string(data), "Hello Go!")
}
