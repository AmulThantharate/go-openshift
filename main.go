package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/health" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"UP"}`)
		return
	}
	fmt.Fprint(w, "🚀 Hello from Go on OpenShift (Day 1 🤓)")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}