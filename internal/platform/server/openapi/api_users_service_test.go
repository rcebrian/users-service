package server

import (
	"context"
	"errors"
	"net/http"
	"testing"

	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/users/mocks"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UsersApiService_CreateUser_BadRequestError(t *testing.T) {
	tests := []struct {
		name        string
		dtoInput    UserDto
		want        ImplResponse
		expectedErr error
	}{
		{
			name:        "error invalid user ID",
			dtoInput:    UserDto{Name: "John", Firstname: "Doe"},
			want:        ImplResponse{Code: 400},
			expectedErr: users.ErrInvalidUserID,
		},
		{
			name:        "error empty user name",
			dtoInput:    UserDto{Name: "", Firstname: "Doe"},
			want:        ImplResponse{Code: 400},
			expectedErr: users.ErrEmptyUserName,
		},
		{
			name:        "error empty user firstname",
			dtoInput:    UserDto{Name: "John", Firstname: ""},
			want:        ImplResponse{Code: 400},
			expectedErr: users.ErrEmptyFirstname,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createServiceMock := new(mocks.CreateUserUseCase)
			findAllServiceMock := new(mocks.FindAllUsersUseCase)
			findByIdServiceMock := new(mocks.FindUserByIdUseCase)

			apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

			createServiceMock.On("Create", mock.Anything, tt.dtoInput.Name, tt.dtoInput.Firstname).
				Return(tt.expectedErr)

			res, err := apiService.CreateUser(context.Background(), tt.dtoInput)

			assert.Equal(t, res.Code, http.StatusBadRequest)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func Test_UsersApiService_CreateUser_InternalServerError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userName, userFirstname := "John", "Doe"

	createServiceMock.On("Create", mock.Anything, userName, userFirstname).
		Return(errors.New("something unexpected happened"))

	dtoInput := UserDto{Name: userName, Firstname: userFirstname}

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.CreateUser(context.Background(), dtoInput)

	assert.Equal(t, res.Code, http.StatusInternalServerError)
	assert.Error(t, err)
}

func Test_UsersApiService_CreateUser_Created(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userName, userFirstname := "John", "Doe"

	createServiceMock.On("Create", mock.Anything, userName, userFirstname).
		Return(nil)

	dtoInput := UserDto{Name: userName, Firstname: userFirstname}

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.CreateUser(context.Background(), dtoInput)

	assert.Equal(t, res.Code, http.StatusCreated)
	assert.Nil(t, err)
}

func Test_UsersApiService_GetAllUsers_NotFoundError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	findAllServiceMock.On("FindAll", mock.Anything).
		Return(nil, users.ErrNotFound)

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetAllUsers(context.Background())

	assert.Equal(t, res.Code, http.StatusNotFound)
	assert.ErrorIs(t, err, users.ErrNotFound)
}

func Test_UsersApiService_GetAllUsers_InternalServerError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	findAllServiceMock.On("FindAll", mock.Anything).
		Return(nil, errors.New("something unexpected happened"))

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetAllUsers(context.Background())

	assert.Equal(t, res.Code, http.StatusInternalServerError)
	assert.Error(t, err)
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

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetAllUsers(context.Background())

	assert.Equal(t, res.Code, http.StatusOK)
	assert.IsType(t, res.Body, GetAllUsers200Response{})
	assert.Nil(t, err)
}

func Test_UsersApiService_GetUserById_NotFoundError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"

	findByIdServiceMock.On("FindById", mock.Anything, userID).
		Return(users.User{}, users.ErrNotFound)

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetUserById(context.Background(), userID)

	assert.Equal(t, res.Code, http.StatusNotFound)
	assert.Error(t, err)
}

func Test_UsersApiService_GetUserById_InternalServerError(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"

	findByIdServiceMock.On("FindById", mock.Anything, userID).
		Return(users.User{}, errors.New("something unexpected happened"))

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetUserById(context.Background(), userID)

	assert.Equal(t, res.Code, http.StatusInternalServerError)
	assert.Error(t, err)
}

func Test_UsersApiService_GetUserById_Ok(t *testing.T) {
	createServiceMock := new(mocks.CreateUserUseCase)
	findAllServiceMock := new(mocks.FindAllUsersUseCase)
	findByIdServiceMock := new(mocks.FindUserByIdUseCase)

	userID, userName, userFirstname := "37a0f027-15e6-47cc-a5d2-64183281087e", "John", "Doe"

	expected, err := users.NewUser(userID, userName, userFirstname)
	require.NoError(t, err)

	findByIdServiceMock.On("FindById", mock.Anything, userID).
		Return(expected, nil)

	apiService := NewUsersApiService(createServiceMock, findAllServiceMock, findByIdServiceMock)

	res, err := apiService.GetUserById(context.Background(), userID)

	assert.Equal(t, res.Code, http.StatusOK)
	assert.IsType(t, res.Body, UserDto{})
	assert.Nil(t, err)
}
