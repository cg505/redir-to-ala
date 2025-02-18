package main

import (
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

const baseURL = "https://archive.archlinux.org"

func main() {
	http.HandleFunc("/", redirectHandler)
	http.ListenAndServe(":8080", nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Split the path into parts
	parts := strings.Split(path, "/")

	// Check if path has at least 4 segments to match /<repo>/os/x86_64/<filename>
	if len(parts) >= 4 && parts[2] == "os" && parts[3] == "x86_64" {
		// Extract the filename from the request path
		filename := filepath.Base(path)

		// Special handling for package database files
		matched, err := regexp.MatchString(`^(.*\.db|.*\.db\.tar\.gz)$`, filename)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if matched {
			repo := parts[1] // Take the first segment as repository name
			newUrl := baseURL + "/repos/last/" + repo + "/os/x86_64/" + filename
			http.Redirect(w, r, newUrl, http.StatusFound)
			return
		}
		// Note: some other files are not handled (repo.files, repo.db.tar.gz.old, etc)

		// For all package files
		newUrl := baseURL + "/packages/.all/" + filename
		http.Redirect(w, r, newUrl, http.StatusFound)
		return
	}

	// Return 404 for non-matching paths
	http.NotFound(w, r)
}
