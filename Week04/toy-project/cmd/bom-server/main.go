package main

import (
	"log"
	"toy"
	v1 "toy/api/bom/v1"
	"toy/internal/di"
	db2 "toy/internal/pkg/db"
	"toy/internal/pkg/server/grpc"
	"toy/internal/service"
)

func init()  {
	db, err := db2.NewDB()
	if err != nil {
		panic(err)
	}
	sql := `create table if not exists bom (id integer not null primary key AUTOINCREMENT, file_name text);`
	_, err = db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func main() {
	app := toy.New()
	ouc, err := di.InitRepo()
	if err != nil {
		panic(err)
	}
	srv := service.NewBomService(ouc.BomUseCase)
	grpcServer := grpc.NewServer("tcp", ":9000")
	v1.RegisterBomServer(grpcServer, srv)
	app.Append(grpcServer)
	if err = app.Run(); err != nil {
		log.Println(err)
	}
}
