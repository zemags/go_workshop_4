package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	MethodNotAllowed = 405
)

func home(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter for combainig http answer and return to user
	// Request pointer for structure thats contain info about current request
	if r.URL.Path != "/" {
		// return 404 if doesnt match to / or other contollers
		http.NotFound(w, r)
		// return to exit from func and do not continue next line
		return
	}
	w.Write([]byte("Hello from 3 chapter"))
}

func showMemo(w http.ResponseWriter, r *http.Request) {
	// get id from URL, and check
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "id is:%d", id)
}

func createMemo(w http.ResponseWriter, r *http.Request) {
	// POST request
	if r.Method != http.MethodPost {
		// add to header 'Allow: POST' user will know
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method GET forbidden", MethodNotAllowed)
		return
	}

	w.Write([]byte("Create Memo"))
}
