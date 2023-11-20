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
