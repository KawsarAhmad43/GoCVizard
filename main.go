package main

import (
	"go-cv-matcher/handlers"
	"net/http"
)

func main() {
	// Serve the upload form at the root
	http.HandleFunc("/", handlers.ShowUploadForm)

	// Handle CV uploads at /upload
	http.HandleFunc("/upload", handlers.UploadCV)

	// Start the server
	port := ":8080"
	println("Server is running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
