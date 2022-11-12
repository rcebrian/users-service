package health

import (
	"api-template/config"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func GetHealth() *health.Service {
	return health.New(
		health.Health{
			ServiceID: config.AppConfig.ServiceID,
			Version:   config.AppConfig.ServiceVersion,
		},
		uptime.System(),
	)
}
