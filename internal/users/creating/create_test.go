package creating

import (
	users "api-template/internal"
	"api-template/internal/platform/storage/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CreateUserUseCase_Create_RepositoryError(t *testing.T) {
	userID, userName, userFirstname := "37a0f027-15e6-47cc-a5d2-64183281087e", "John", "Doe"

	_, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.Anything).
		Return(errors.New("something unexpected happened"))

	userService := NewCreatingService(userRepositoryMock)

	err = userService.Create(context.Background(), userName, userFirstname)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CreateUserUseCase_Create_Succeed(t *testing.T) {
	userName, userFirstname := "John", "Doe"

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("users.User")).
		Return(nil)

	userService := NewCreatingService(userRepositoryMock)

	err := userService.Create(context.Background(), userName, userFirstname)

	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
