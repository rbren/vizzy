package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const baseDir = "test/e2e/cases/"

//go:embed iframe.html
var templateFS embed.FS

func findIndexHTMLFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Base(path) == "index.html" {
			path = strings.TrimPrefix(path, baseDir)
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err // Return any error that occurred during the walk.
	}

	return files, nil // Return the collected 'index.html' paths.
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := 0
	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			log.Fatalf("Error parsing page: %v", err)
		}
	}
	indexes, err := findIndexHTMLFiles(baseDir)
	if err != nil {
		log.Fatalf("Error finding index.html files: %v", err)
	}
	if page < 0 || page >= len(indexes) {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	// Define your data structure that matches the placeholders in your HTML template.
	data := struct {
		IframeURL string
		PrevURL   string
		NextURL   string
	}{
		IframeURL: indexes[page],
		PrevURL:   "/home?page=" + strconv.Itoa(page-1), // or "" if not applicable
		NextURL:   "/home?page=" + strconv.Itoa(page+1), // or "" if not applicable
	}
	if page == 0 {
		data.PrevURL = ""
	}
	if page == len(indexes)-1 {
		data.NextURL = ""
	}

	// Parse the embedded template.
	tmpl, err := template.ParseFS(templateFS, "iframe.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute the template and write the output to the response writer.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// noCache is a middleware that sets cache-control headers.
func noCache(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")                       // HTTP/1.0.
		w.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT") // Proxies.
		h.ServeHTTP(w, r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./test/e2e/cases/"))
	assets := http.FileServer(http.Dir("./app/src/assets/"))

	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.Handle("/", noCache(fs))
	http.HandleFunc("/home", homeHandler)

	log.Println("Listening on :3333...")
	err := http.ListenAndServe("0.0.0.0:3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
