package main

import (
	"fmt"
	"net/http"

	h "sample.com/handlers"
)

func main() {
	http.HandleFunc("/status/", h.StatusHandler)

	fmt.Printf("Starting sample service at :%d\n", 3000)
	http.ListenAndServe(":3000", nil)
}
