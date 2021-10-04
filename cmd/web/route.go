package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//I have wrapped access to the custom errors log and handlers
func (app *application) routes() *chi.Mux{
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Get("/", app.home)
	mux.Post("/process", app.processPForm)
	mux.Get("/blog", app.blog)
	mux.Get("/blog/{UserId}", app.Delete)
	//staticServer := http.FileServer(http.Dir("./ui/static"))
	//mux.Handle("/static/", http.StripPrefix("/static", staticServer))

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./ui/static/"))
	FileServer(mux, "/static", filesDir)

	return mux
}
// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}