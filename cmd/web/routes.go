package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/ping", app.ping)
	mux.HandleFunc("/snippet", app.getSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir(app.config.StaticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
