package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"snippetbox.doinby.net/internal/models"
)

// Define an application struct to hold the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	// Make SnippetModel object available to handler functions
	snippets *models.SnippetModel
}

func main() {
	// Define a command line flag with a default of 3000
	addr := flag.String("addr", ":3000", "HTTP network access")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Defer a call to db.Close() so the connection pool is closed before main() exits
	defer db.Close()

	// Initialise new instance of app which contains dependencies
	// Now we can use our application struct's methods as handler func
	// in Routes section
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,

		// Initialise models.SnippetModel instance and add it to application dependencies
		snippets: &models.SnippetModel{DB: db},
	}

	// Initialize a new http.Server struct,
	// set the Addr and Handler fields so
	// server uses same network address and routes as before
	// and set the ErrorLog field to use custom errorLog logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // from ./routes.go
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Function which uses Go's sql.Open() and returns sql.Open() connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
