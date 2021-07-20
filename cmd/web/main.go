package main

import (
	"log"
	"net/http"
)

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
