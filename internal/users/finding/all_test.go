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

func Test_AllUseCase_FindAll_RepositoryError(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).
		Return(nil, errors.New("something unexpected happened"))

	userService := NewFindAllUsersUseCase(userRepositoryMock)

	_, err := userService.FindAll(context.Background())

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_AllUseCase_FindAll_NotFoundError(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).
		Return(nil, sql.ErrNoRows)

	userService := NewFindAllUsersUseCase(userRepositoryMock)

	_, err := userService.FindAll(context.Background())

	userRepositoryMock.AssertExpectations(t)
	assert.ErrorIs(t, err, users.ErrNotFound)
}

func Test_AllUseCase_FindAll_FoundUsers(t *testing.T) {
	userID, userName, userFirstname := "37a0f027-15e6-47cc-a5d2-64183281087e", "John", "Doe"

	user, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	expected := []users.User{user}

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("FindAll", mock.Anything).
		Return(expected, nil)

	userService := NewFindAllUsersUseCase(userRepositoryMock)

	got, err := userService.FindAll(context.Background())
	require.NoError(t, err)

	userRepositoryMock.AssertExpectations(t)
	assert.Equal(t, expected, got)
}
