package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// Serve static files from the web directory
	fs := http.FileServer(http.Dir("web"))
	
	// Handle all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// For wasm files, set the correct MIME type
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}
		
		// Serve the file
		fs.ServeHTTP(w, r)
	})
	
	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
