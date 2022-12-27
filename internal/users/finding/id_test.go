package finding

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/platform/storage/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ByIdUseCase_FindById_RepositoryError(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindById", mock.Anything, userID).
		Return(users.User{}, errors.New("something unexpected happened"))

	userService := NewFindUserByIdUseCase(userRepositoryMock)

	_, err := userService.FindById(context.Background(), userID)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_ByIdUseCase_FindById_NotFoundError(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindById", mock.Anything, userID).
		Return(users.User{}, sql.ErrNoRows)

	userService := NewFindUserByIdUseCase(userRepositoryMock)

	_, err := userService.FindById(context.Background(), userID)

	userRepositoryMock.AssertExpectations(t)
	assert.ErrorIs(t, err, users.ErrNotFound)
}

func Test_ByIdUseCase_FindById_FoundUser(t *testing.T) {
	userID, userName, userFirstname := "37a0f027-15e6-47cc-a5d2-64183281087e", "John", "Doe"

	expected, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindById", mock.Anything, userID).
		Return(expected, nil)

	userService := NewFindUserByIdUseCase(userRepositoryMock)

	got, err := userService.FindById(context.Background(), userID)
	require.NoError(t, err)

	userRepositoryMock.AssertExpectations(t)
	assert.Equal(t, expected, got)
}
