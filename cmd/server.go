package main

import (
	"log"
	"net/http"

	"github.com/eidosam/gecho/pkg/gecho"
)

func main() {
	http.HandleFunc("/", gecho.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
