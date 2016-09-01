package main

import (
	"net/http"

	h "sample.com/handlers"
)

func main() {
	http.HandleFunc("/status/", h.StatusHandler)
	http.ListenAndServe(":3000", nil)
}
