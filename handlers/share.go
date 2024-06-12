package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		slog.Error("HTTPHandler error", "err", err, "path", r.URL.Path)
		http.Error(w, err.Error(), 500)
	}
}

func render(w http.ResponseWriter, r *http.Request, templ templ.Component) error {
	return templ.Render(r.Context(), w)
}
