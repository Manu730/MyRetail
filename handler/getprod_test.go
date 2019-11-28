package hndlr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/products/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("{id}", "13860428")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Result":"SUCCESS","RespCode":"NONE","RespDescription":"NONE","product":{"id":13860428,"name":"The Big Lebowski (Blu-ray)","current_price":{"value":25,"currency_code":"USD"}}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
