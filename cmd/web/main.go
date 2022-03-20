package main

import (
	"crypto/tls"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/rhodeon/prettylog"
	"github.com/rhodeon/sniphub/cmd/web/internal/server"
	"github.com/rhodeon/sniphub/cmd/web/internal/templates"
	"github.com/rhodeon/sniphub/pkg/mailer"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	sqlDriver "github.com/go-sql-driver/mysql"
	"github.com/rhodeon/sniphub/pkg/models/mysql"
)

func main() {
	// configure sessionFlags
	sessionFlags := flags{}
	sessionFlags.parse()
	err := sessionFlags.validate()
	if err != nil {
		prettylog.FatalError(err)
	}

	// initiate database connection
	dbConfig := &sqlDriver.Config{
		DBName:               sessionFlags.sqlDb,
		User:                 sessionFlags.sqlUser,
		Passwd:               sessionFlags.sqlPassword,
		Addr:                 sessionFlags.sqlAddr,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dsn := dbConfig.FormatDSN()
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

	sessionMailer := mailer.New(sessionFlags.smtpHost, sessionFlags.smtpPort, sessionFlags.smtpUser, sessionFlags.smtpPass)

	app := server.Application{
		TemplateCache:  templateCache,
		SessionManager: sessionManager,
		Mailer:         sessionMailer,
		Snips:          &mysql.SnipController{Db: db},
		Users:          &mysql.UserController{Db: db},
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         sessionFlags.addr,
		Handler:      app.RouteHandler(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	prettylog.InfoF("Starting server on %s", sessionFlags.addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	prettylog.Error(err.Error())
}
