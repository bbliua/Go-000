// +build wireinject

package di

import (
	"github.com/google/wire"
	"toy/internal/biz"
	"toy/internal/data"
	db2 "toy/internal/pkg/db"
)
//go:generate wire
func InitRepo() (*Repo, error)  {
	panic(wire.Build(db2.NewDB, data.Provider, biz.NewBomUseCase, NewRepo))
}
