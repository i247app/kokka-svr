package app

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/i247app/gex"
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/app/routes"
	"kokka.com/kokka/internal/app/services"
	"kokka.com/kokka/internal/handlers/http/middleware"
	"kokka.com/kokka/internal/shared/config"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/response"
)

func NewFromEnv(envPath string) (*App, error) {
	// Load configuration
	env, err := config.NewEnv(envPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// Build app resource
	hostConfig := gex.HostConfig{
		ServerHost: env.HostConfig.ServerHost,
		ServerPort: env.HostConfig.ServerPort,
	}
	if env.HostConfig.HttpsCertFile != nil {
		hostConfig.HttpsCertFile = *env.HostConfig.HttpsCertFile
	}
	if env.HostConfig.HttpsKeyFile != nil {
		hostConfig.HttpsKeyFile = *env.HostConfig.HttpsKeyFile
	}
	resources := resources.AppResource{
		Env:        env,
		HostConfig: hostConfig,
	}

	app := NewApp(&resources)
	if err := app.Init(); err != nil {
		return nil, fmt.Errorf("failed to init app: %w", err)
	}

	routes.SetUpHttpRoutes(app.Server, &resources, app.Services)

	return app, nil
}

func (a *App) Init() error {
	services, err := services.SetupServiceContainer(a.Resource)
	if err != nil {
		return fmt.Errorf("failed to setup services: %w", err)
	}
	a.Services = services

	defaultRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("route not found"), status.NOT_FOUND)
	}
	a.Server = gex.NewServer(a.Resource.HostConfig, defaultRouteHandler)

	// Register middlewares
	a.setupMiddleware(a.Server, services)

	// Setup jobs

	// Setup shutdown hooks
	a.setupShutdownHooks(a.Server, services)

	// Reload sessions

	return nil
}

func (a *App) Start() error {
	return a.Server.Start()
}

func (a *App) setupShutdownHooks(gexServer *gex.Server, _ *services.ServiceContainer) {

}

// Setup middlewares
func (a *App) setupMiddleware(gexSvr *gex.Server, services *services.ServiceContainer) {
	middlewares := []gex.Middleware{
		// Start-->
		middleware.LoggerMiddleware(a.Resource.Env.LogFile),
		middleware.LogRequestMiddleware,
		// -->End
	}

	slices.Reverse(middlewares) // Reverse the middleware order so that the first middleware in the slice is the first to run
	for _, middleware := range middlewares {
		gexSvr.RegisterMiddleware(middleware)
	}

	gexSvr.SetupServerCORS()
}

func (a *App) Close() error {
	return nil
}
