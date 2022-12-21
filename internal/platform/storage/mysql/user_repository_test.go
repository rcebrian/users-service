package mysql

import (
	users "api-template/internal"
	"api-template/pkg/logger"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	logger.SetOutput(io.Discard)
}

func Test_UserRepository_Save_RepositoryError(t *testing.T) {
	userID, userName, userFirstname := "02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe"
	user, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO user (id, name, firstname) VALUES (?, ?, ?)").
		WithArgs(userID, userName, userFirstname).
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(db)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_UserRepository_Save_Succeed(t *testing.T) {
	userID, userName, userFirstname := "02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe"
	user, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO user (id, name, firstname) VALUES (?, ?, ?)").
		WithArgs(userID, userName, userFirstname).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(db)

	err = repo.Save(context.Background(), user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_UserRepository_FindById_RepositoryError(t *testing.T) {
	userID, _, _ := "02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT * FROM user WHERE id = ? LIMIT 1").
		WithArgs(userID).
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(db)

	_, err = repo.FindById(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_UserRepository_FindById_Success(t *testing.T) {
	userID, userName, userFirstname := "02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMockRows := sqlMock.NewRows([]string{"id", "name", "firstname"}).
		AddRow(userID, userName, userFirstname)

	sqlMock.ExpectQuery("SELECT * FROM user WHERE id = ? LIMIT 1").
		WithArgs(userID).
		WillReturnRows(sqlMockRows)

	expectedUser, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	repo := NewUserRepository(db)

	actual, _ := repo.FindById(context.Background(), userID)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, expectedUser, actual)
}

func Test_UserRepository_FindAll_RepositoryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"SELECT * FROM user").
		WillReturnError(errors.New("something-failed"))

	repo := NewUserRepository(db)

	_, err = repo.FindAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_UserRepository_FindAll_Success(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMockRows := sqlMock.NewRows([]string{"id", "name", "firstname"}).
		AddRow("02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe").
		AddRow("29ed61bf-1f9b-40bb-bffd-377c6367260d", "Evita", "Peachy")

	sqlMock.ExpectQuery("SELECT * FROM user").WillReturnRows(sqlMockRows)

	expectedUser1, _ := users.NewUser("02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe")
	expectedUser2, _ := users.NewUser("29ed61bf-1f9b-40bb-bffd-377c6367260d", "Evita", "Peachy")

	expectedUsers := []users.User{expectedUser1, expectedUser2}

	repo := NewUserRepository(db)

	actual, _ := repo.FindAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Equal(t, expectedUsers, actual)
}
