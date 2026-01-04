package resources

import (
	"github.com/i247app/gex"
	"kokka.com/kokka/internal/shared/config"
)

type AppResource struct {
	Env        *config.Env
	HostConfig gex.HostConfig
}
