package health

import (
	"api-template/configs"
	"api-template/internal/platform/storage/mysql"
	"database/sql"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/uptime"
)

func GetHealth(client *sql.DB) *health.Service {
	return health.New(
		health.Health{
			ServiceID: configs.ServiceConfig.ServiceID,
			Version:   configs.ServiceConfig.ServiceVersion,
		},
		uptime.System(),
		mysql.Health("mysql", client, configs.MySqlConfig.Timeout, configs.MySqlConfig.Threshold),
	)
}
