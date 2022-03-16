package main

import (
	"database/sql"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/rhodeon/prettylog"
	server2 "github.com/rhodeon/sniphub/cmd/web/internal/server"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rhodeon/sniphub/pkg/models/mysql"
)

func main() {
	// configure sessionFlags
	sessionFlags := flags{}
	sessionFlags.parse()

	// initiate database connection
	dsn := sessionFlags.sqlUser + ":" + sessionFlags.sqlPassword + "@/" + sessionFlags.sqlDb + "?parseTime=true"
	db, err := openDb(dsn)
	if err != nil {
		prettylog.FatalError(err.Error())
	}
	defer db.Close()

	// initialize a new template cache
	templateCache, err := templates.NewTemplateCache("./ui/html")
	if err != nil {
		prettylog.FatalError(err)
	}

	// initialize session manager
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	app := server2.Application{
		TemplateCache:  templateCache,
		SessionManager: sessionManager,
		Snips:          &mysql.SnipController{Db: db},
		Users:          &mysql.UserController{Db: db},
	}

	server := &http.Server{
		Addr:         sessionFlags.addr,
		Handler:      app.RouteHandler(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	prettylog.InfoF("Starting server on %s", sessionFlags.addr)
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
