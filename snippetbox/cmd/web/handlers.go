package main

import (
	"errors"
	"fmt"
	//"html/template"
	"net/http"
	"snippetbox.doinby.net/internal/models"
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

	// Show latest snippets
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// Initialize a slice function containing paths to templates.
	// It's important to note that the file containing our BASE tmpl must be
	// the *first* file in the slice.
	//files := []string{
	//	"./ui/html/base.tmpl.html",
	//	"./ui/html/partials/nav.tmpl.html",
	//	"./ui/html/pages/home.tmpl.html",
	//}

	// Parse HTML template into variable
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.errorLog.Print(err.Error())
	//	app.serverError(w, err)
	//	return
	//}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.errorLog.Print(err.Error())
	//	app.serverError(w, err)
	//}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Check if query "id" is valid
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use SnippetModel object's Get method to retrieve the data for a specific
	// record based on its ID. If no match were found, return 404 response
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w) // DB error
		} else {
			app.serverError(w, err) // Server error
		}
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check if method is valid
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
