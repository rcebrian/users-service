package health

import (
	"api-template/config"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func GetHealth() *health.Service {
	return health.New(
		health.Health{
			Version:   config.AppConfig.AppVersion,
			ReleaseID: config.AppConfig.AppVersion,
		},
		uptime.System(),
	)
}
