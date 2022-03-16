package main

import "database/sql"

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
