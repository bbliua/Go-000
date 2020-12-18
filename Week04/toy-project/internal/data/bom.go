package data

import (
	"database/sql"
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"toy/internal/biz"
)
var Provider = wire.NewSet(NewBomRepo)
var _ biz.BomRepo = (biz.BomRepo)(nil)

func NewBomRepo(db *sql.DB) biz.BomRepo  {
	return &bomRepo{db: db}
}

type bomRepo struct {
	db *sql.DB
}

func (br *bomRepo) SaveBom(b *biz.Bom) error {
	// 数据持久
	tx, err := br.db.Begin()
	if err != nil {
		return errors.Wrap(err, "db error")
	}
	sql :="insert into bom(file_name) values(?)"
	stmt, err := tx.Prepare(sql)
	_, err = stmt.Exec(b.FileName)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("sql: %s, arg: %s", sql, b.FileName))
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "transaction commit error")
	}
	return err
}
