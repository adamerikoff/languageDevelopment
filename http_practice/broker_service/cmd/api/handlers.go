package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse = jsonResponse{
		Error:   false,
		Message: "My Message from Handler.go",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}
