package handlers

import (
	"net/http"
)

func shouldReturn(r *http.Request) string {
	var header = r.Header

	if header.Get("Accept") == "application/json" {
		return "json"
	}

	if header.Get("Hx-Request") == "true" {
		return "partial"
	}

	return "full"
}
