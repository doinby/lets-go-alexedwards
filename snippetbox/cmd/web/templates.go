package main

import "snippetbox.doinby.net/internal/models"

// Define a templateData type to act as the holding structure
// for any dynamic data that we want to pass into HTML template
type templateData struct {
	Snippet *models.Snippet
}
