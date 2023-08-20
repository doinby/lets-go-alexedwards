package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Add (app *application) as signature of the home handler, so it is defined
// as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if valid URL
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice function containing paths to templates.
	// It's important to note that the file containing our BASE tmpl must be
	// the *first* file in the slice.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Parse HTML template into variable
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.serverError(w, err)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Check if query "id" is valid
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check if method is valid
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create new snippet..."))
}
