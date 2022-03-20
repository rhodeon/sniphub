package main

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

// flags houses the flags for the current session.
type flags struct {
	addr string

	sqlDb       string
	sqlUser     string
	sqlPassword string
	sqlAddr     string

	smtpHost string
	smtpPort int
	smtpUser string
	smtpPass string
}

// parse registers the properties of f with the current flag contents and parses them afterwards.
func (f *flags) parse() {
	flag.StringVar(&f.addr, "addr", ":"+port(), "HTTP network address")

	flag.StringVar(&f.sqlDb, "sqlDb", "sniphub_dev", "SQL database name")
	flag.StringVar(&f.sqlUser, "sqlUser", "web", "SQL user name")
	flag.StringVar(&f.sqlPassword, "sqlPassword", "password", "SQL database password")
	flag.StringVar(&f.sqlAddr, "sqlAddr", "localhost:3306", "SQL host and port")

	flag.StringVar(&f.smtpHost, "smtpHost", f.defSmtpHost(), "SMTP hostname")
	flag.IntVar(&f.smtpPort, "smtpPort", f.defSmtpPort(), "SMTP port")
	flag.StringVar(&f.smtpUser, "smtpUser", f.defSmtpUser(), "SMTP username")
	flag.StringVar(&f.smtpPass, "smtpPass", f.defSmtpPass(), "SMTP password")

	flag.Parse()
}

// validate flags and returns an error if a required flag isn't set.
func (f *flags) validate() error {
	if f.smtpHost == "" {
		return errors.New("env: the smtpHost flag is required")
	}

	if f.smtpUser == "" {
		return errors.New("env: the smtpUser flag is required")
	}

	return nil
}

func port() string {
	const defaultPort = "4000"
	if port, exists := os.LookupEnv("PORT"); exists {
		return port
	}
	return defaultPort
}

func (f *flags) defSmtpHost() string {
	if host, exists := os.LookupEnv("SMTP_HOST"); exists {
		return host
	}
	return ""
}

func (f *flags) defSmtpPort() int {
	defaultPort := 587

	if portEnv, exists := os.LookupEnv("SMTP_PORT"); exists {
		port, err := strconv.Atoi(portEnv)
		if err == nil {
			return port
		}
	}
	return defaultPort
}

func (f *flags) defSmtpUser() string {
	if user, exists := os.LookupEnv("SMTP_USER"); exists {
		return user
	}
	return ""
}

func (f *flags) defSmtpPass() string {
	if password, exists := os.LookupEnv("SMTP_PASS"); exists {
		return password
	}
	return ""
}
