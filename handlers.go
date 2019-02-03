package main

import (
	"fmt"
	"net/http"
)

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called. Hi")
}
