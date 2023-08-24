package models

import (
	"database/sql"
	"errors"
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
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO snippets (title, content, created, expires)
    		 VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content and expiry values for the placeholder parameters. This
	// method returns a sql.Result type, which contains some basic
	// information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// Get returns a snippet correspond to its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE
			 expires > UTC_TIMESTAMP() AND id = ?`

	// Use QueryRow() method on the connection pool to execute SQL statement,
	// passing in untrusted id var and returns a sql.Row object which holds
	// the result from the DB
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct
	s := &Snippet{}

	// Use row.Scan() to copy the values to each field in sql.Row to the
	// corresponding field in the Snippet struct. Note that the argument
	// to row.Scan are *pointers* and the number of argument must match
	// with Snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// Check whether the error caused by query not found using
		// errors.Is() function from errors package
		if errors.Is(err, sql.ErrNoRows) {
			// Run ErrNoRecord func from ./errors.go
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// If everything went OK, return Snippet object
	return s, nil
}

// Latest returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
