package main

import (
	"bytes"
	"encoding/json"
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

func UnmarshalBody(buffer *bytes.Buffer, value any) {
	data, _ := io.ReadAll(buffer)

	_ = json.Unmarshal(data, value)
}

func TestGetAllBooksEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/book", nil)
	rec := recordRequest(req)

	assert.Equal(t, 200, rec.Code)

	var actual []Book
	UnmarshalBody(rec.Body, &actual)
	assert.Equal(t, books, actual)
}

func TestGetBookByIdEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/book/2", nil)
	rec := recordRequest(req)

	assert.Equal(t, 200, rec.Code)

	var book Book
	UnmarshalBody(rec.Body, &book)

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

func TestAddBookEndpoint(t *testing.T) {
	book := Book{6, "Go Power", 50}
	jsonValue, _ := json.Marshal(book)

	req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))

	rec := recordRequest(req)

	assert.Equal(t, 201, rec.Code)

	actual := Book{}
	UnmarshalBody(rec.Body, &actual)

	assert.Equal(t, book, actual)
}

func TestAddBookEndpointShouldFailForInvalidInput(t *testing.T) {
	jsonValue := []byte("{\"Title\":\"Go Power\"}")

	req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))

	rec := recordRequest(req)

	assert.Equal(t, 400, rec.Code)
}

func TestUpdateBookHandler(t *testing.T) {
	book := Book{1, "Awesome Go", 110}
	jsonValue, _ := json.Marshal(book)

	req := httptest.NewRequest("PUT", "/book", bytes.NewBuffer(jsonValue))
	rec := recordRequest(req)

	assert.Equal(t, http.StatusOK, rec.Code)

	actual := Book{}
	UnmarshalBody(rec.Body, &actual)

	assert.Equal(t, book, actual)
}

func TestUpdateBookHandlerShouldFailIfBookNotFound(t *testing.T) {
	book := Book{1000, "Awesome Go", 110}
	jsonValue, _ := json.Marshal(book)

	req := httptest.NewRequest("PUT", "/book", bytes.NewBuffer(jsonValue))
	rec := recordRequest(req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateBookHandlerShouldFailForInvalidInput(t *testing.T) {
	jsonValue := []byte("{\"Title\":\"Go Power\"}")

	req := httptest.NewRequest("PUT", "/book", bytes.NewBuffer(jsonValue))
	rec := recordRequest(req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
