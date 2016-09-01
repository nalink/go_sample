package handlers

import (
	"net/http"
	"io"

	"sample.com/util"
)

func StatusHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == util.HttpPost {
		http.Error(w, "Invalid call", http.StatusMethodNotAllowed)
		io.WriteString(w, util.StatusFail)
		return
	}

	io.WriteString(w, util.StatusOK)
}
