package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	status := response.Status

	fmt.Println(statusCode)
	fmt.Println(status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Husein", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	status := response.Status

	fmt.Println(statusCode)
	fmt.Println(status)
	fmt.Println(string(body))
}
