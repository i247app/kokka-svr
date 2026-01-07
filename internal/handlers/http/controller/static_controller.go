package controller

import (
	"net/http"
	"os"
	"path/filepath"
)

type StaticController struct {
	basePath string
}

func NewStaticController(basePath string) *StaticController {
	return &StaticController{
		basePath: basePath,
	}
}

func (c *StaticController) ServeFile(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fullPath := filepath.Join(c.basePath, filePath)

		// Check if file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// Set appropriate content type
		ext := filepath.Ext(filePath)
		switch ext {
		case ".html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		case ".css":
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}

		// Serve the file
		http.ServeFile(w, r, fullPath)
	}
}
