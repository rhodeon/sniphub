package mysql

import (
	"database/sql"
	"os"
	"testing"
)

// newTestDb generates a database for helping function.
// It creates and returns a function for dropping concerned tables whenever called.
func newTestDb(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	db, err := sql.Open("mysql", "test_web:password@/test_sniphub?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	// create tables
	execScript(t, db, "./testdata/setup.sql")

	// return function to drop tables
	return db, func() {
		execScript(t, db, "./testdata/teardown.sql")
		defer db.Close()
	}
}

// execScript is a helper function to execute SQL commands
// in the file at the given scriptPath
func execScript(t *testing.T, db *sql.DB, scriptPath string) {
	t.Helper()

	script, err := os.ReadFile(scriptPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}
}
