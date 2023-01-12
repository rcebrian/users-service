package health

import "github.com/nelkinda/health-go"

// NewHealthService creates a health service that implements RFC healthcheck standard
func NewHealthService(ID, version string, checks ...health.ChecksProvider) *health.Service {
	return health.New(
		health.Health{
			ServiceID: ID,
			Version:   version,
		},
		checks...)
}
