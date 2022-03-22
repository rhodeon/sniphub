package server

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

// Env houses the configurations for the current session.
type Env struct {
	Addr    string
	BaseUrl string

	SqlDb       string
	SqlUser     string
	SqlPassword string
	SqlAddr     string

	SmtpHost string
	SmtpPort int
	SmtpUser string
	SmtpPass string
}

// Parse registers the properties of f with the current flag contents and parses them afterwards.
func (e *Env) Parse() {
	flag.StringVar(&e.Addr, "addr", ":"+e.port(), "HTTP network address")
	flag.StringVar(&e.BaseUrl, "baseUrl", e.defBaseUrl(), "base URL of server")

	flag.StringVar(&e.SqlDb, "sqlDb", "sniphub_dev", "SQL database name")
	flag.StringVar(&e.SqlUser, "sqlUser", "web", "SQL user name")
	flag.StringVar(&e.SqlPassword, "sqlPassword", "password", "SQL database password")
	flag.StringVar(&e.SqlAddr, "sqlAddr", "localhost:3306", "SQL host and port")

	flag.StringVar(&e.SmtpHost, "smtpHost", e.defSmtpHost(), "SMTP hostname")
	flag.IntVar(&e.SmtpPort, "smtpPort", e.defSmtpPort(), "SMTP port")
	flag.StringVar(&e.SmtpUser, "smtpUser", e.defSmtpUser(), "SMTP username")
	flag.StringVar(&e.SmtpPass, "smtpPass", e.defSmtpPass(), "SMTP password")

	flag.Parse()
}

// Validate Env and returns an error if a required flag isn't set.
func (e *Env) Validate() error {
	if e.SmtpHost == "" {
		return errors.New("env: the smtpHost flag is required")
	}

	if e.SmtpUser == "" {
		return errors.New("env: the smtpUser flag is required")
	}

	return nil
}

func (e *Env) port() string {
	const defaultPort = "4000"
	if port, exists := os.LookupEnv("PORT"); exists {
		return port
	}
	return defaultPort
}

func (e *Env) defBaseUrl() string {
	defaultHost := "https://localhost:4000"
	if baseUrl, exists := os.LookupEnv("BASE_URL"); exists {
		return baseUrl
	}
	return defaultHost
}

func (e *Env) defSmtpHost() string {
	if host, exists := os.LookupEnv("SMTP_HOST"); exists {
		return host
	}
	return ""
}

func (e *Env) defSmtpPort() int {
	defaultPort := 587

	if portEnv, exists := os.LookupEnv("SMTP_PORT"); exists {
		port, err := strconv.Atoi(portEnv)
		if err == nil {
			return port
		}
	}
	return defaultPort
}

func (e *Env) defSmtpUser() string {
	if user, exists := os.LookupEnv("SMTP_USER"); exists {
		return user
	}
	return ""
}

func (e *Env) defSmtpPass() string {
	if password, exists := os.LookupEnv("SMTP_PASS"); exists {
		return password
	}
	return ""
}
