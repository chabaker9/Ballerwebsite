package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// TemplateData represents the data to pass to templates.
type TemplateData struct {
	Method        string
	URL           string
	Protocol      string
	Host          string
	Header        http.Header
	RemoteAddr    string
	RequestURI    string
	ContentLength int64
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Load HTML templates
	tmpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))
	router.SetHTMLTemplate(tmpl)

	// Serve static files
	router.Static("/static", "static")

	// Define a route for the index page
	router.GET("/", func(c *gin.Context) {
		data := TemplateData{
			Method:        c.Request.Method,
			URL:           c.Request.URL.String(),
			Protocol:      c.Request.Proto,
			Host:          c.Request.Host,
			Header:        c.Request.Header,
			RemoteAddr:    c.Request.RemoteAddr,
			RequestURI:    c.Request.RequestURI,
			ContentLength: c.Request.ContentLength,
		}
		c.HTML(http.StatusOK, "index.html", data)
	})

	// Define a route for all other pages
	router.NoRoute(func(c *gin.Context) {
		path := strings.Trim(c.Request.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		if tmpl.Lookup(path) != nil {
			data := TemplateData{
				Method:        c.Request.Method,
				URL:           c.Request.URL.String(),
				Protocol:      c.Request.Proto,
				Host:          c.Request.Host,
				Header:        c.Request.Header,
				RemoteAddr:    c.Request.RemoteAddr,
				RequestURI:    c.Request.RequestURI,
				ContentLength: c.Request.ContentLength,
			}
			c.HTML(http.StatusOK, path, data)
		} else {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
		}
	})

	// Start server
	router.Run(":8080")
}

// package main

// import (
// 	"html/template"
// 	"net/http"
// 	"path/filepath"
// 	"strings"
// )

// // TemplateData represents the data to pass to templates.
// type TemplateData struct {
// 	Method        string
// 	URL           string
// 	Protocol      string
// 	Host          string
// 	Header        http.Header
// 	RemoteAddr    string
// 	RequestURI    string
// 	ContentLength int64
// 	// Add other fields as needed
// }

// func main() {
// 	// Parse templates
// 	tmpl, err := template.ParseGlob(filepath.Join("templates", "*.html"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Serve static files
// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// Handler for all template requests
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Extract requested template name from URL path
// 		requestedFile := strings.Trim(r.URL.Path, "/")
// 		if requestedFile == "" {
// 			requestedFile = "index.html" // Default to index.html if no file is specified
// 		}

// 		// Prepare template data
// 		data := TemplateData{
// 			Method:        r.Method,
// 			URL:           r.URL.String(),
// 			Protocol:      r.Proto,
// 			Host:          r.Host,
// 			Header:        r.Header,
// 			RemoteAddr:    r.RemoteAddr,
// 			RequestURI:    r.RequestURI,
// 			ContentLength: r.ContentLength,
// 			// Set other data fields as needed
// 		}

// 		// Check if the requested template exists
// 		if tmpl.Lookup(requestedFile) != nil {
// 			// Execute the template with data
// 			err := tmpl.ExecuteTemplate(w, requestedFile, data)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 		} else {
// 			// Render custom 404 page if template is not found
// 			w.WriteHeader(http.StatusNotFound)
// 			err := tmpl.ExecuteTemplate(w, "404.html", nil)
// 			if err != nil {
// 				// Fallback error handling if 404 template is also missing
// 				http.Error(w, "404 Not Found", http.StatusNotFound)
// 			}
// 		}
// 	})

// 	// Start server
// 	http.ListenAndServe(":8080", nil)
// }
