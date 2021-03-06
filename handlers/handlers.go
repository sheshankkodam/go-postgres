package handlers

import (
	"fmt"
	"github.com/sheshankkodam/go-postgres/config"
	"net/http"
)

// HelloHandler takes a GET parameter "name" and responds // with Hello <name>! in plaintext
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(config.Version)))
}
