package helpers

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Success bool   `json:"success"`
	Length  *int   `json:"length,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Err     *string `json:"err,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}
