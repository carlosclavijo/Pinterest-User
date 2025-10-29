package helpers

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	resp := Response[string]{
		Success: true,
		Length:  ptr(1),
		Data:    "hello",
	}

	WriteJSON(rr, http.StatusOK, resp)

	result := rr.Result()
	defer result.Body.Close()

	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	assert.Contains(t, string(body), `"success":true`)
	assert.Contains(t, string(body), `"data":"hello"`)
}

func TestWriteJSON_EncodingError(t *testing.T) {
	rr := httptest.NewRecorder()
	badData := make(chan int)

	WriteJSON(rr, http.StatusOK, badData)

	result := rr.Result()
	defer result.Body.Close()

	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	assert.Contains(t, string(body), "Failed to encode response")
}

func ptr[T any](v T) *T {
	return &v
}
