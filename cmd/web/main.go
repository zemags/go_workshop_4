package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// application - for dependency injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	addr := flag.String("addr", ":8080", "Network address")
	flag.Parse()

	// add logger
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

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

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("start listening server: %v", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
