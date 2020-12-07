package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"week.02/internal/dao"
	"week.02/internal/di"
	"week.02/internal/e"
)

var iniDb int

func init() {
	flag.IntVar(&iniDb, "iniDb", 0, "set argument --iniDb=1 if neened")
	flag.Parse()
	if iniDb == 1 {
		db, err := dao.NewDB()
		if err != nil {
			panic(err)
		}
		sql := `create table bom (id integer not null primary key, file_name text);delete from bom`
		_, err = db.Exec(sql)
		if err != nil {
			panic(err)
		}
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}
		stmt, err := tx.Prepare("insert into bom(id, file_name) values(?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		for i := 0; i < 10; i++ {
			_, err = stmt.Exec(i, fmt.Sprintf("file name xx %03d", i))
			if err != nil {
				panic(err)
			}
		}
		tx.Commit()
	}
}

func main() {
	app, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	defer closeFunc()
	ctx := context.Background()
	// rsp, err := app.FetchBom(ctx, 1)
	rsp, err := app.FetchBom(ctx, -1)
	if err != nil {
		// handle error
		// 此处简单打印处理
		if errors.Is(err, e.ErrRecordNotFound) {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%v\n", err)
		}
		os.Exit(1)
	}
	fmt.Printf("id: %d, file name: %s", rsp.Id, rsp.FileName)
}
