package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called. Hi")

	json.NewEncoder(w).Encode("Welcome!")
}
