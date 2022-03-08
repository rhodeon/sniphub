package main

import "flag"

// flags houses the flags for the current session.
type flags struct {
	addr        string
	sqlDb       string
	sqlUser     string
	sqlPassword string
}

// parse registers the properties of f with the current flag contents and parses them afterwards.
func (f *flags) parse() {
	flag.StringVar(&f.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&f.sqlDb, "sqlDb", "sniphub", "SQL database name")
	flag.StringVar(&f.sqlUser, "sqlUser", "web", "SQL user name")
	flag.StringVar(&f.sqlPassword, "sqlPassword", "password", "SQL database password")
	flag.Parse()
}