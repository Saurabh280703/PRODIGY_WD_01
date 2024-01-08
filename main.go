package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the requested file path based on the URL path
	filePath := r.URL.Path[1:] // Remove the leading slash

	// Serve the appropriate content type for the requested file
	contentType := ""
	switch {
	case filePath == "":
		// Serve index.html as the default page
		filePath = "index.html"
		contentType = "text/html"
	case filePath == "style.css":
		contentType = "text/css"
	case filePath == "script.js":
		contentType = "application/javascript"
	case strings.HasSuffix(filePath, ".png"):
		contentType = "image/png"
	case strings.HasSuffix(filePath, ".jpg"):
		contentType = "image/jpg"
	case strings.HasSuffix(filePath, ".svg"):
		contentType = "image/svg+xml"
	case strings.HasSuffix(filePath, ".svg"):
		contentType = "image/svg+xml"
	case strings.HasSuffix(filePath, ".svg"):
		contentType = "image/svg+xml"

	default:
		http.NotFound(w, r)
		return
	}

	// Open the requested file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set the Content-Type header
	w.Header().Set("Content-Type", contentType)

	// Copy the content of the file to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not copy file content to response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := 8087 // Set the port number you want to use
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
