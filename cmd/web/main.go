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

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("start listening server: %v", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
