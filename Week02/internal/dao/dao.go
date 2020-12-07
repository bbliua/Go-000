package dao

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"week.02/internal/model"
)

var Provider = wire.NewSet(New, NewDB)

// Dao dao interface
type Dao interface {
	Bom(ctx context.Context, id int64) (bom *model.Bom, err error)
	Close()
}

type dao struct {
	db *sql.DB
}

// New new Dao
func New(db *sql.DB) (d Dao, cf func(), err error) {
	d = &dao{
		db: db,
	}
	cf = d.Close
	return
}

func (d *dao) Bom(ctx context.Context, id int64) (bom *model.Bom, err error) {
	return d.RawBom(ctx, id)
}

func (d *dao) Close() {
	d.db.Close()
}
