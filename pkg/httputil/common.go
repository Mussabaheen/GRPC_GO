package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	// HTTP error code
	Code int `json:"code" example:"400"`

	// HTTP error message
	Message string `json:"message" example:"status bad request"`
}

// RespondWithError writes to a http.ResponseWriter with an error string
func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code == http.StatusInternalServerError {
		fmt.Println(message)
	}
	RespondWithJSON(w, code, HTTPError{Code: code, Message: message})
}

// RespondWithError writes to a http.ResponseWriter with an error object
func RespondWithErrorObject(w http.ResponseWriter, code int, err error) {
	if code == http.StatusInternalServerError {
		fmt.Println(err)
	}
	RespondWithJSON(w, code, HTTPError{Code: code, Message: err.Error()})
}

//RespondWithJSON writes to a http.ResponseWriter with a json payload
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "could not marshal json response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
