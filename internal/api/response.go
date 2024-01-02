package api

import (
	"net/http"

	"github.com/goccy/go-json"
)

type (
	// The expected structure of each http response
	// returned by the dataserver.
	StdResponse struct {
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	// It should be used only in http responses when a struct is
	// not available and at the same time will only be used once.
	Map map[string]any
)

const (
	Error   string = "Error"
	Success string = "Success"
)

// Adds a header with 'application/json' and the status code to the provided http.ResponseWriter.
//
// On error, throws an error with an internal server error status code and body.
func Response(w http.ResponseWriter, statusCode int, v any) {
	content, err := json.Marshal(v)
	if err != nil {
		SendInternalServerError(w)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
}

// Shortcut for bad request response.
func SendNotFound(w http.ResponseWriter, message ...string) {
	responseShortcut(w, http.StatusNotFound, "Not Found", message...)
}

// Shortcut for internal server error response.
func SendInternalServerError(w http.ResponseWriter, message ...string) {
	responseShortcut(w, http.StatusInternalServerError, "Internal Server Error", message...)
}

func responseShortcut(w http.ResponseWriter, statusCode int, defaultLabel string, label ...string) {
	if len(label) != 0 {
		defaultLabel = label[0]
	}

	Response(w, statusCode, StdResponse{Message: defaultLabel})
}
