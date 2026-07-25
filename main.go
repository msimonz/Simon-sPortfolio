// Command portfolio serves Simón Márquez's portfolio, and can also render it to
// a folder of static files for hosting anywhere.
//
//	go run .              → http://localhost:8080
//	go run . -port 3000   → another port
//	go run . -build dist  → writes dist/index.html plus assets, then exits
//
// Standard library only: net/http, html/template, embed. No dependencies.
package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"simonmarquez/portfolio/internal/content"
)

//go:embed web/templates/*.html
var templateFS embed.FS

//go:embed web/static
var staticFS embed.FS

func main() {
	port := flag.Int("port", 8080, "puerto del servidor")
	buildDir := flag.String("build", "", "renderiza el sitio estático en esta carpeta y termina")
	flag.Parse()

	tmpl, err := parseTemplates()
	if err != nil {
		log.Fatalf("plantillas: %v", err)
	}

	site := content.Load()

	if *buildDir != "" {
		if err := build(tmpl, site, *buildDir); err != nil {
			log.Fatalf("build: %v", err)
		}
		fmt.Printf("Sitio estático escrito en %s\n", *buildDir)
		return
	}

	if err := serve(tmpl, site, *port); err != nil {
		log.Fatalf("servidor: %v", err)
	}
}

// parseTemplates loads the HTML templates. There are no template helpers on
// purpose: everything the page shows is already shaped correctly in
// internal/content, so the markup is plain field access and ranges.
func parseTemplates() (*template.Template, error) {
	return template.ParseFS(templateFS, "web/templates/*.html")
}

// serve runs the site over HTTP, with the embedded assets mounted at /static.
func serve(tmpl *template.Template, site content.Site, port int) error {
	assets, err := fs.Sub(staticFS, "web/static")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(assets))))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "index.html", site); err != nil {
			log.Printf("render: %v", err)
		}
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Portafolio de %s en http://localhost%s\n", site.Profile.Given, addr)
	return http.ListenAndServe(addr, mux)
}

// build writes the site to disk as plain files, so it can be hosted on GitHub
// Pages, Netlify, an OCI bucket, or anything else that serves static content.
func build(tmpl *template.Template, site content.Site, dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	page, err := os.Create(filepath.Join(dir, "index.html"))
	if err != nil {
		return err
	}
	defer page.Close()

	if err := tmpl.ExecuteTemplate(page, "index.html", site); err != nil {
		return err
	}

	return copyStatic(dir)
}

// copyStatic mirrors the embedded asset tree into the build directory.
func copyStatic(dir string) error {
	return fs.WalkDir(staticFS, "web/static", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel("web", path)
		if err != nil {
			return err
		}
		target := filepath.Join(dir, rel)

		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		src, err := staticFS.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(target)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})
}
