// +build wireinject

package di

import (
	"github.com/google/wire"
	"week.02/internal/dao"
	"week.02/internal/service"
)

//go:generate wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, NewApp))
}
