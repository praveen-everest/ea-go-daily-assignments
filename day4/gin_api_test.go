package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func recordRequest(req *http.Request) *httptest.ResponseRecorder {
	r := setUpRouter()

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	return rec
}

func TestGetAllBooksEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/book", nil)
	rec := recordRequest(req)

	assert.Equal(t, 200, rec.Code)

	body := rec.Body
	data, _ := io.ReadAll(body)
	var actual []Book
	_ = json.Unmarshal(data, &actual)
	assert.Equal(t, books, actual)
}

func TestGetBookByIdEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/book/2", nil)
	rec := recordRequest(req)

	assert.Equal(t, 200, rec.Code)

	body := rec.Body
	data, _ := io.ReadAll(body)
	fmt.Println(string(data))
	var book Book
	_ = json.Unmarshal(data, &book)

	expected := Book{2, "Learn Go!", 90}
	assert.Equal(t, expected, book)
}

func TestGetBookIdShouldReturn404IfBookNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/book/1000", nil)
	rec := recordRequest(req)

	assert.Equal(t, 404, rec.Code)

	body := rec.Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, "{\"message\":\"book with id 1000 not found\"}", string(data))
}
