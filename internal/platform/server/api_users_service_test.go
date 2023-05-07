package server

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UsersApiService_GetAllUsers_NotFoundError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	findAllServiceMock.On("FindAll", mock.Anything).
		Return(nil, users.ErrNotFound)

	apiService := NewUsersApiServer(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetAllUsers(context.Background(), GetAllUsersRequestObject{})

	assert.ErrorIs(t, err, users.ErrNotFound)
	assert.Equal(t, res, GetAllUsers404JSONResponse{})
}

func Test_UsersApiService_GetAllUsers_InternalServerError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	findAllServiceMock.On("FindAll", mock.Anything).
		Return(nil, errors.New("something unexpected happened"))

	apiService := NewUsersApiServer(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetAllUsers(context.Background(), GetAllUsersRequestObject{})

	assert.Error(t, err)
	assert.Equal(t, res, GetAllUsers500JSONResponse{})
}

func Test_UsersApiService_GetAllUsers_Ok(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userID, userName, userFirstname := "37a0f027-15e6-47cc-a5d2-64183281087e", "John", "Doe"

	user, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	expected := []users.User{user}

	findAllServiceMock.On("FindAll", mock.Anything).
		Return(expected, nil)

	apiService := NewUsersApiServer(createServiceMock, findAllServiceMock, findByIdServiceMock)

	got, err := apiService.GetAllUsers(context.Background(), GetAllUsersRequestObject{})

	assert.Nil(t, err)
	assert.IsType(t, got, GetAllUsers200JSONResponse{})
}
