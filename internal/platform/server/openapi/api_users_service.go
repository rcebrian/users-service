package server

import (
	users "api-template/internal"
	"context"
	"net/http"
)

// UsersApiService is a users that implements the logic for the UsersApiServicer
// This users should implement the business logic for every endpoint for the UsersApi API.
// Include any external packages or services that will be required by these users.
type UsersApiService struct {
	userRepository users.UserRepository
}

// NewUsersApiService creates a default api users
func NewUsersApiService(repository users.UserRepository) UsersApiServicer {
	return &UsersApiService{
		userRepository: repository,
	}
}

// CreateUser - Save user into data storage
func (s *UsersApiService) CreateUser(ctx context.Context, dto UserDto) (ImplResponse, error) {
	user, err := users.NewUser(dto.Id, dto.Name, dto.Firstname)
	if err != nil {
		return Response(http.StatusBadRequest, nil), err
	}

	err = s.userRepository.Save(ctx, user)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	return Response(http.StatusCreated, nil), err
}

// GetAllUsers - Get all users
func (s *UsersApiService) GetAllUsers(ctx context.Context) (ImplResponse, error) {
	all, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	var usersDto = make([]UserDto, len(all))

	for i := range all {
		usersDto[i].Id = all[i].ID().String()
		usersDto[i].Name = all[i].Name().String()
		usersDto[i].Firstname = all[i].Firstname().String()
	}

	return Response(http.StatusOK, usersDto), nil
}

// GetUserById - Get user by id
func (s *UsersApiService) GetUserById(ctx context.Context, userId string) (ImplResponse, error) {
	user, err := s.userRepository.FindById(ctx, userId)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	if user.ID().String() == "" {
		return Response(http.StatusNotFound, nil), err
	}

	return Response(http.StatusOK, UserDto{
		Id:        user.ID().String(),
		Name:      user.Name().String(),
		Firstname: user.Firstname().String(),
	}), nil
}
