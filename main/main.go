package main

import (
	"fmt"
	"net/http"
)

// Your function
func TestHtml() string {
	return "<h1> Zainab is Pretty </h1>"
}

// Handler function
func handler(w http.ResponseWriter, r *http.Request) {
	html := TestHtml()
	fmt.Fprint(w, html)
}

func main() {
	// Route
	http.HandleFunc("/", handler)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}