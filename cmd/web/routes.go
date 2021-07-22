package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// create new router mux and register our controlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/memo", app.showMemo)
	mux.HandleFunc("/memo/create", app.createMemo)

	// init fileserver to work with static files by http requests
	fileServer := http.FileServer(safeFileSystem{http.Dir("./ui/static/")})
	// mux.Handle("/static", http.NotFoundHandler())

	// register handler for request with /static/, and remove /static from path
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
