package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostProductSuccess(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          1,
		Name:        "Some name",
		Description: "Some description",
		Price:       1.5,
		Currency:    "EUR",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,
		"{\"id\":1,\"name\":\"Some name\",\"description\":\"Some description\",\"price\":1.5,\"currency\":\"EUR\"}",
		w.Body.String())
}

func TestPostProductWithoutID(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          0,
		Name:        "Some name",
		Description: "Some description",
		Price:       1.5,
		Currency:    "EUR",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,
		"{\"code\":400,\"message\":\"Key: 'Product.ID' Error:Field validation for 'ID' failed on the 'required' tag\",\"status\":\"Bad Request\"}",
		w.Body.String())
}

func TestPostProductWithoutName(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          1,
		Name:        "",
		Description: "Some description",
		Price:       1.5,
		Currency:    "EUR",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,
		"{\"code\":400,\"message\":\"Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'required' tag\",\"status\":\"Bad Request\"}",
		w.Body.String())
}

func TestPostProductMaxLimitForName(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          1,
		Name:        "this is too long a name to assign",
		Description: "Some description",
		Price:       1.5,
		Currency:    "EUR",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,
		"{\"code\":400,\"message\":\"Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'max' tag\",\"status\":\"Bad Request\"}",
		w.Body.String())
}

func TestPostProductPriceLessThanZero(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          1,
		Name:        "Some name",
		Description: "Some description",
		Price:       -1.5,
		Currency:    "EUR",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,
		"{\"code\":400,\"message\":\"Key: 'Product.Price' Error:Field validation for 'Price' failed on the 'min' tag\",\"status\":\"Bad Request\"}",
		w.Body.String())
}

func TestPostProductInvalidCurrency(t *testing.T) {
	r := Setup()
	w := httptest.NewRecorder()
	prd := Product{
		ID:          1,
		Name:        "Some name",
		Description: "Some description",
		Price:       1.5,
		Currency:    "FAKE",
	}
	body, err := json.Marshal(prd)

	assert.Equal(t, nil, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t,
		"{\"code\":400,\"message\":\"Key: 'Product.Currency' Error:Field validation for 'Currency' failed on the 'iso4217' tag\",\"status\":\"Bad Request\"}",
		w.Body.String())
}
