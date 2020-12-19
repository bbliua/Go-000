// +build wireinject

package di

import (
	"github.com/google/wire"
	"toy/internal/biz"
	"toy/internal/data"
	db2 "toy/internal/pkg/db"
)
//go:generate wire
func InitUseCase() (*UseCase, error)  {
	panic(wire.Build(db2.NewDB, data.Provider, biz.NewBomUseCase, NewUseCase))
}
