package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter for combainig http answer and return to user
	// Request pointer for structure thats contain info about current request
	w.Write([]byte("hello from 2 chapter"))
}

func main() {
	// create new router mux and register controler home
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("start listening server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
