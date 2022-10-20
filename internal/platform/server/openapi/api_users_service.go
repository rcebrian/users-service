package server

import (
	model "api-template/internal"
	"context"
	"net/http"
)

// UsersApiService is a service that implements the logic for the UsersApiServicer
// This service should implement the business logic for every endpoint for the UsersApi API.
// Include any external packages or services that will be required by this service.
type UsersApiService struct {
	userRepository model.UserRepository
}

// NewUsersApiService creates a default api service
func NewUsersApiService(repository model.UserRepository) UsersApiServicer {
	return &UsersApiService{
		userRepository: repository,
	}
}

// CreateUser - Save user into data storage
func (s *UsersApiService) CreateUser(ctx context.Context, dto UserDto) (ImplResponse, error) {
	err := s.userRepository.Save(ctx, model.NewUser(dto.Id, dto.Name, dto.Firstname))
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	return Response(http.StatusCreated, nil), err
}

// GetAllUsers - Get all users
func (s *UsersApiService) GetAllUsers(ctx context.Context) (ImplResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	var usersDto = make([]UserDto, len(users))

	for i := range users {
		usersDto[i].Id = users[i].ID()
		usersDto[i].Name = users[i].Name()
		usersDto[i].Firstname = users[i].Firstname()
	}

	return Response(http.StatusOK, usersDto), nil
}

// GetUserById - Get user by id
func (s *UsersApiService) GetUserById(ctx context.Context, userId int32) (ImplResponse, error) {
	user, err := s.userRepository.FindById(ctx, string(userId))
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	if user.ID() == "" {
		return Response(http.StatusNotFound, nil), err
	}

	return Response(http.StatusOK, UserDto{
		Id:        user.ID(),
		Name:      user.Name(),
		Firstname: user.Firstname(),
	}), nil
}
