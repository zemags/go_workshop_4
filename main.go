package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter for combainig http answer and return to user
	// Request pointer for structure thats contain info about current request
	if r.URL.Path != "/" {
		// return 404 if doesnt match to / or other contollers
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from 2 chapter"))
}

func showMemo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show Memo"))
}

func createMemo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Memo"))
}

func main() {
	// create new router mux and register our controlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/memo", showMemo)
	mux.HandleFunc("/memo/create", createMemo)

	log.Println("start listening server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
