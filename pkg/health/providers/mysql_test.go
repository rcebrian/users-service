package providers

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nelkinda/health-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	componentID                     = "test-component"
	timeout           time.Duration = 5000
	affectedEndpoints               = []string{"/foo", "/fuu"}
)

func Test_Mysql_HealthChecks_ShouldPass(t *testing.T) {
	client, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	mysqlHealthMock := NewMysqlProvider(componentID, affectedEndpoints, client, timeout, 200)

	sqlMock.ExpectPing().WillDelayFor(150 * time.Millisecond)

	got := mysqlHealthMock.HealthChecks()
	actual := got["mysql:responseTime"][0]

	assert.Equal(t, health.Pass, actual.Status)
}

func Test_Mysql_HealthChecks_ShouldWarn(t *testing.T) {
	client, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	mysqlHealthMock := NewMysqlProvider(componentID, affectedEndpoints, client, timeout, 150)

	sqlMock.ExpectPing().WillDelayFor(200 * time.Millisecond)

	got := mysqlHealthMock.HealthChecks()
	actual := got["mysql:responseTime"][0]

	assert.Equal(t, health.Warn, actual.Status)
	assert.Equal(t, affectedEndpoints, actual.AffectedEndpoints)
}

func Test_Mysql_HealthChecks_ShouldFail(t *testing.T) {
	mockErr := errors.New("ping failed")
	client, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	mysqlHealthMock := NewMysqlProvider(componentID, affectedEndpoints, client, timeout, 200)

	sqlMock.ExpectPing().WillDelayFor(150 * time.Millisecond).WillReturnError(mockErr)

	got := mysqlHealthMock.HealthChecks()
	actual := got["mysql:responseTime"][0]

	assert.Equal(t, health.Fail, actual.Status)
	assert.Error(t, mockErr, actual.Output)
	assert.Equal(t, affectedEndpoints, actual.AffectedEndpoints)
}

func Test_Mysql_HealthChecks_ShouldFailConnection(t *testing.T) {
	client, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	_ = client.Close()

	mysqlHealthMock := NewMysqlProvider(componentID, affectedEndpoints, client, timeout, 200)

	sqlMock.ExpectPing().WillDelayFor(150 * time.Millisecond)

	got := mysqlHealthMock.HealthChecks()
	actual := got["mysql:responseTime"][0]

	assert.Equal(t, health.Fail, actual.Status)
	assert.Equal(t, affectedEndpoints, actual.AffectedEndpoints)
}

func Test_Mysql_AuthorizeHealth_IsTrue(t *testing.T) {
	mysqlHealthMock := NewMysqlProvider(componentID, affectedEndpoints, nil, timeout, 150)

	got := mysqlHealthMock.AuthorizeHealth(nil)

	assert.True(t, got)
}
