package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel defines type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert adds new snippet to DB
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// Get returns a snippet correspond to its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Latest returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
