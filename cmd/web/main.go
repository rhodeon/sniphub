package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rhodeon/sniphub/pkg/models/mysql"
	"github.com/rhodeon/sniphub/pkg/prettylog"
)

func main() {
	// configure flags
	addr, sqlDb, sqlUser, sqlPassword := parseFlags()

	// initiate database connection
	dsn := *sqlUser + ":" + *sqlPassword + "@/" + *sqlDb + "?parseTime=true"
	db, err := openDb(dsn)
	if err != nil {
		prettylog.FatalError(err.Error())
	}
	defer db.Close()

	// initialize a new template cache
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		prettylog.FatalError(err)
	}

	// initialize session manager
	sessionManager := scs.New()

	app := application{
		templateCache:  templateCache,
		sessionManager: sessionManager,
		snips:          &mysql.SnipController{Db: db},
		users:          &mysql.UserController{Db: db},
	}

	server := &http.Server{
		Addr:         *addr,
		Handler:      app.routesHandler(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	prettylog.InfoF("Starting server on %s", *addr)
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
