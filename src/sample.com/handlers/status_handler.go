package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"sample.com/util"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == util.HttpPost {
		http.Error(w, "Invalid call", http.StatusMethodNotAllowed)
		io.WriteString(w, util.StatusFail)
		return
	}

	fmt.Printf("[%d] StatusHandler called\n", time.Now().Unix())
	io.WriteString(w, util.StatusOK)
}
