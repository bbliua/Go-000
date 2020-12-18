package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)
var dbFile = "bom.db"

// NewDB return db sql.DB
func NewDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		err = errors.Wrap(err, "db connect error")
	}
	// defer db.Close()
	return
}
