package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/zemags/go_workshop_4/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// ResponseWriter - for combainig http answer and return to user
	// Request - pointer for structure thats contain info about current request
	if r.URL.Path != "/" {
		// return 404 if doesnt match to / or other contollers
		app.notFound(w)
		// 'return' to exit from func and do not continue next lines
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) showMemo(w http.ResponseWriter, r *http.Request) {
	// get id from URL (../memo?id=1), and check
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	m, err := app.memos.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", m)
}

func (app *application) createMemo(w http.ResponseWriter, r *http.Request) {
	// POST request need
	if r.Method != http.MethodPost {
		// add to header 'Allow: POST' and user will know
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := app.memos.Insert("Melon", "Five melons", "3")
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/memo?id=%d", id), http.StatusSeeOther)
}
