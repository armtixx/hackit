package handlers

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

// Get

// Add/

// Update/

// Delete/
