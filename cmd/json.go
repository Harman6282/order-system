package main

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type successResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *any   `json:"data,omitempty"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := errorResponse{
		Success: false,
		Message: message,
	}

	return json.NewEncoder(w).Encode(res)
}

func writeJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := successResponse{
		Success: true,
		Message: message,
		Data:    &data,
	}

	return json.NewEncoder(w).Encode(res)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxbytes := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxbytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}
