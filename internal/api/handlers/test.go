package handlers

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
