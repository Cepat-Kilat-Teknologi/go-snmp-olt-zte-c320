package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequestJSONResponse(t *testing.T) {
	// Inisialisasi ResponseWriter dan Request
	rr := httptest.NewRecorder()

	// Contoh respons yang ingin Anda kirim
	response := WebResponse{
		Code:   200,
		Status: "OK",
		Data:   map[string]string{"key": "value"},
	}

	// Panggil fungsi SendJSONResponse
	SendJSONResponse(rr, http.StatusOK, response)

	// Periksa kode status respons
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusOK)
	}

	// Periksa tipe konten
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa respons JSON
	var decodedResponse WebResponse
	err := json.NewDecoder(rr.Body).Decode(&decodedResponse)
	if err != nil {
		t.Errorf("Gagal mendekode respons JSON: %v", err)
	}

}

func TestErrorBadRequestBos(t *testing.T) {
	// Inisialisasi ResponseWriter
	rr := httptest.NewRecorder()

	// Contoh error
	err := errors.New("Bad Request")

	// Panggil fungsi ErrorBadRequest
	ErrorBadRequest(rr, err)

	// Periksa kode status respons
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code tidak sesuai: got %v want %v", status, http.StatusBadRequest)
	}

	// Periksa tipe konten
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type tidak sesuai: got %v want %v", contentType, expectedContentType)
	}

	// Periksa respons JSON
	var decodedResponse ErrorResponse
	err = json.NewDecoder(rr.Body).Decode(&decodedResponse)
	if err != nil {
		t.Errorf("Gagal mendekode respons JSON: %v", err)
	}

	expectedResponse := ErrorResponse{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: "Bad Request",
	}

	if decodedResponse != expectedResponse {
		t.Errorf("Respons JSON tidak sesuai: got %+v want %+v", decodedResponse, expectedResponse)
	}
}
