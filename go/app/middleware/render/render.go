package render

import (
	"context"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/unrolled/render"
)

type contextKey string

// const ...
const (
	renderContextKey contextKey = "render/render_key"
)

// Builder ...
type Builder struct {
}

// MiddleWare ...
func (b *Builder) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		renderEngine := render.New(render.Options{
			IndentJSON:    (r.Form.Get("pretty") == "true"),
			IsDevelopment: true,
		})
		r = r.WithContext(context.WithValue(r.Context(), renderContextKey, renderEngine))
		next.ServeHTTP(w, r)
	})
}

// GetRenderer ...
func GetRenderer(r *http.Request) *render.Render {
	return r.Context().Value(renderContextKey).(*render.Render)
}

// JSON ...
func JSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render := GetRenderer(r)
	render.JSON(w, status, data)
}

// HTML ...
func HTML(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	tpl, err := pongo2.DefaultSet.FromFile(file)
	if err != nil {
		panic(err)
	}
	context := pongo2.Context{
		"dummy_key": "dummy value",
	}
	if data != nil {
		for k, v := range data.(map[string]interface{}) {
			context[k] = v
		}
	}
	tpl.ExecuteWriter(context, w)
}
