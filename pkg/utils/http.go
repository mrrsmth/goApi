package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type JSONError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// RespondJSON sends a JSON response to the client.
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := JSONResponse{
		Status: status,
		Data:   data,
	}
	json.NewEncoder(w).Encode(response)
}

// RespondError sends an error response to the client.
func RespondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := JSONError{
		Status:  status,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// ParseJSON parses a JSON request body into a struct.
func ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
