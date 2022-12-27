package health

import (
	"database/sql"

	"github.com/rcebrian/users-service/configs"
	"github.com/rcebrian/users-service/internal/platform/storage/mysql"

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
