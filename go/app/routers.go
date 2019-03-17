package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/dongri/web-app-on-docker/go/app/handlers"
	"github.com/dongri/web-app-on-docker/go/app/middleware/render"
)

var defaultMiddlewares = []func(next http.Handler) http.Handler{
	middleware.Logger,
	middleware.Recoverer,
	middleware.StripSlashes,
}

// NewRootHandler returns a root handler.
func NewRootHandler() http.Handler {
	r := chi.NewRouter()

	// setup middlewares
	renderBuilder := &render.Builder{}
	middlewares := append(defaultMiddlewares, renderBuilder.MiddleWare)

	r.Use(middlewares...)

	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	r.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("./public/img"))))
	r.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	r.Handle("/manifest.json", http.FileServer(http.Dir("./public/pwa")))
	r.Handle("/service-worker.js", http.FileServer(http.Dir("./public/pwa")))

	r.Get("/", handlers.IndexHandler)

	return r
}
