package users_api_config

import "github.com/rcebrian/users-service/configs"

const ServiceID = "users-api-service"

func ConfigureServer() error {
	err := configs.ConfigureLogger()

	return err
}
