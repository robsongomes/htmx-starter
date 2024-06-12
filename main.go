package main

import (
	"net/http"

	"github.com/robsongomes/htmx-starter/handlers"
)

func main() {

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("GET /", handlers.HTTPHandler(handlers.HomeHandler))

	http.ListenAndServe(":3000", nil)
}
