package providers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/nelkinda/health-go"
)

type mysqldb struct {
	componentID string
	client      *sql.DB
	timeout     time.Duration
	threshold   time.Duration
}

// NewMysqlProvider create a new mysql healthcheck
func NewMysqlProvider(componentID string, client *sql.DB, timeout, threshold time.Duration) health.ChecksProvider {
	return &mysqldb{
		componentID: componentID,
		client:      client,
		timeout:     timeout * time.Millisecond,
		threshold:   threshold * time.Millisecond,
	}
}

// HealthChecks implements mysql checks
func (m *mysqldb) HealthChecks() map[string][]health.Checks {
	start := time.Now().Local()
	startTime := start.Format(time.RFC3339Nano)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	var checks = health.Checks{
		ComponentID:   m.componentID,
		Time:          startTime,
		ComponentType: "datastore",
	}

	conn, err := m.client.Conn(ctxTimeout)
	if err != nil {
		checks.Output = err.Error()
		checks.Status = health.Fail

		return map[string][]health.Checks{"mysql:responseTime": {checks}}
	}

	err = conn.PingContext(ctxTimeout)
	_ = conn.Close()

	if err != nil {
		checks.Output = err.Error()
		checks.Status = health.Fail

		return map[string][]health.Checks{"mysql:responseTime": {checks}}
	}

	end := time.Now().Local()
	responseTime := end.Sub(start)
	checks.ObservedValue = responseTime.Nanoseconds()
	checks.ObservedUnit = "ns"

	if responseTime.Nanoseconds() > m.threshold.Nanoseconds() {
		checks.Status = health.Warn
	} else {
		checks.Status = health.Pass
	}

	return map[string][]health.Checks{"mysql:responseTime": {checks}}
}

// AuthorizeHealth enables auth as default
func (m *mysqldb) AuthorizeHealth(_ *http.Request) bool {
	return true
}
