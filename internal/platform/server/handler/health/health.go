package health

import (
	"api-template/config"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func GetHealth() *health.Service {
	return health.New(
		health.Health{
			ServiceID: config.ServiceConfig.ServiceID,
			Version:   config.ServiceConfig.ServiceVersion,
		},
		uptime.System(),
	)
}
