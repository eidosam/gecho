package gecho

import (
	"net/http"
)

// Handler :: HTTP Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
}
