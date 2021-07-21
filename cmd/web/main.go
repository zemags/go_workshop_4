package main

import (
	"log"
	"net/http"
	"path/filepath"
)

type safeFileSystem struct {
	fs http.FileSystem
}

// Open - call every time when http.FileServer received request
func (sfs safeFileSystem) Open(path string) (http.File, error) {
	f, err := sfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := sfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}

func main() {
	// create new router mux and register our controlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/memo", showMemo)
	mux.HandleFunc("/memo/create", createMemo)

	// init fileserver to work with static files by http requests
	fileServer := http.FileServer(safeFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())

	// register handler for request with /static/, and remove /static from path
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("start listening server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
