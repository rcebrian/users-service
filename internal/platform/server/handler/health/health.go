package health

import (
	"api-template/config"
	"api-template/internal/platform/storage/mysql"
	"database/sql"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func GetHealth(client *sql.DB) *health.Service {
	return health.New(
		health.Health{
			ServiceID: config.ServiceConfig.ServiceID,
			Version:   config.ServiceConfig.ServiceVersion,
		},
		uptime.System(),
		mysql.Health("mysql", client, config.MySqlConfig.Timeout, config.MySqlConfig.Threshold),
	)
}
