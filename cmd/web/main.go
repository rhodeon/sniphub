package main

import (
	"database/sql"
	"flag"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rhodeon/sniphub/pkg/prettylog"
)

func main() {
	// configure flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	sqlDb := flag.String("sqlDb", "sniphub", "SQL database name")
	sqlUser := flag.String("sqlUser", "web", "SQL user name")
	sqlPass := flag.String("sqlPass", "password", "SQL database password")
	flag.Parse()

	app := application{}

	// initiate database connection
	dsn := *sqlUser + ":" + *sqlPass + "@/" + *sqlDb + "?parseTime=true"
	db, err := openDb(dsn)
	if err != nil {
		prettylog.FatalError(err.Error())
	}
	defer db.Close()

	// start server
	prettylog.InfoF("Starting server on %s", *addr)
	err = http.ListenAndServe(*addr, app.routesHandler())
	prettylog.Error(err.Error())
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	// establish a connection to the database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
