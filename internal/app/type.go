package app

import (
	"github.com/i247app/gex"
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/app/services"
)

type App struct {
	Server   *gex.Server
	Services *services.ServiceContainer

	Resource *resources.AppResource
}

func NewApp(resource *resources.AppResource) *App {
	return &App{
		Resource: resource,
	}
}
