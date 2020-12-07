package di

import "week.02/internal/service"

type App struct {
	*service.Service
}

func NewApp(svc *service.Service) (app *App, err error) {
	app = &App{svc}
	return
}
