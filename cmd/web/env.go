package main

import (
	"flag"
	"os"
)

// flags houses the flags for the current session.
type flags struct {
	addr        string
	sqlDb       string
	sqlUser     string
	sqlPassword string
	sqlAddr     string
}

// parse registers the properties of f with the current flag contents and parses them afterwards.
func (f *flags) parse() {
	flag.StringVar(&f.addr, "addr", ":"+port(), "HTTP network address")
	flag.StringVar(&f.sqlDb, "sqlDb", "sniphub_dev", "SQL database name")
	flag.StringVar(&f.sqlUser, "sqlUser", "web", "SQL user name")
	flag.StringVar(&f.sqlPassword, "sqlPassword", "password", "SQL database password")
	flag.StringVar(&f.sqlAddr, "sqlAddr", "localhost:3306", "SQL host and port")
	flag.Parse()
}

const defaultPort = "4000"

func port() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = defaultPort
	}
	return port
}
