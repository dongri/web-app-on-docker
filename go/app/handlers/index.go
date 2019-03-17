package handlers

import (
	"net/http"

	"github.com/dongri/web-app-on-docker/go/app/middleware/render"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	const view = "app/views/index.html"
	data := map[string]interface{}{
		"message": "hello world",
	}
	render.HTML(w, r, view, data)
}
