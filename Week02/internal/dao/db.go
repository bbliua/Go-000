package dao

import (
	"context"
	"database/sql"
	"flag"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"week.02/internal/e"
	"week.02/internal/model"
)

var dbFile string

func init() {
	flag.StringVar(&dbFile, "db", "", "default sqlite db file")
}

// NewDB return db sql.DB
func NewDB() (db *sql.DB, err error) {
	if dbFile == "" {
		err = errors.Wrap(err, "please set sqlite db")
	}
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		err = errors.Wrap(err, "db connect error")
	}
	// defer db.Close()
	return
}

func (d *dao) RawBom(ctx context.Context, id int64) (bom *model.Bom, err error) {
	bom = &model.Bom{}
	err = d.db.QueryRowContext(ctx, "select file_name from bom where id = ?", id).Scan(&bom.FileName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.Wrap(e.ErrRecordNotFound, "404 Not Found")
		} else {
			err = errors.Wrap(e.ErrQueryFail, "sqlite3 query error")
		}
		return
	}

	bom.ID = id
	return
}
