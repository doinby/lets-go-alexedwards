package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Log error msg and stack trace to errorLog
// and response with error 500
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// errorLog defined in main.go
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Output error code and description for user to see
// Example: "404 - Bad Request"
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Outputs a clientError func with specific 404 status
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
