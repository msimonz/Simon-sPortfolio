// Command portfolio serves Simón Márquez's portfolio, and can also render it to
// a folder of static files for hosting anywhere.
//
//	go run .              → http://localhost:8080
//	go run . -port 3000   → another port
//	go run . -build dist  → writes dist/index.html, dist/es/index.html plus assets, then exits
//
// English lives at the root and Spanish under /es/; a switch in the top-right
// corner of the page moves between them.
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
	"net/url"
	"os"
	"path/filepath"
	"strings"

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

	if *buildDir != "" {
		if err := build(tmpl, *buildDir); err != nil {
			log.Fatalf("build: %v", err)
		}
		fmt.Printf("Sitio estático escrito en %s\n", *buildDir)
		return
	}

	if err := serve(tmpl, *port); err != nil {
		log.Fatalf("servidor: %v", err)
	}
}

// parseTemplates loads the HTML templates.
func parseTemplates() (*template.Template, error) {
	return template.New("").Funcs(helpers()).ParseFS(templateFS, "web/templates/*.html")
}

// helpers holds the one thing the template cannot do on its own. Everything
// else is plain field access, because internal/content already shapes the data.
func helpers() template.FuncMap {
	return template.FuncMap{"href": href}
}

// href marks a link as safe to emit.
//
// html/template only trusts http, https and mailto, and silently rewrites
// anything else to "#ZgotmplZ" — which is what happened to the tel: link on the
// phone button. Marking the value as template.URL skips that check, so this
// re-does the check itself and only lets through the four schemes the site
// actually uses. Anything else is neutralised, exactly as before.
func href(raw string) template.URL {
	parsed, err := url.Parse(raw)
	if err != nil {
		return "#"
	}

	switch strings.ToLower(parsed.Scheme) {
	case "http", "https", "mailto", "tel":
		return template.URL(raw)
	default:
		return "#"
	}
}

// serve runs the site over HTTP, with the embedded assets mounted at /static and
// one page per language: / for English, /es/ for Spanish.
func serve(tmpl *template.Template, port int) error {
	assets, err := fs.Sub(staticFS, "web/static")
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(assets))))
	for _, lang := range content.Langs {
		site := content.Load(lang)
		mux.HandleFunc("GET /"+lang.Path()+"{$}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := tmpl.ExecuteTemplate(w, "index.html", site); err != nil {
				log.Printf("render: %v", err)
			}
		})
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Portafolio en http://localhost%s (español en /es/)\n", addr)
	return http.ListenAndServe(addr, mux)
}

// build writes the site to disk as plain files, so it can be hosted on GitHub
// Pages, Netlify, an OCI bucket, or anything else that serves static content.
func build(tmpl *template.Template, dir string) error {
	for _, lang := range content.Langs {
		if err := buildPage(tmpl, content.Load(lang), filepath.Join(dir, lang.Path())); err != nil {
			return err
		}
	}

	return copyStatic(dir)
}

// buildPage renders one language's page as dir/index.html.
func buildPage(tmpl *template.Template, site content.Site, dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	page, err := os.Create(filepath.Join(dir, "index.html"))
	if err != nil {
		return err
	}
	defer page.Close()

	return tmpl.ExecuteTemplate(page, "index.html", site)
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
