package server

import (
	users "api-template/internal"
	"api-template/internal/users/creating"
	"api-template/internal/users/finding"
	"context"
	"errors"
	"net/http"
)

// UsersApiService is a users that implements the logic for the UsersApiServicer
// This users should implement the business logic for every endpoint for the UsersApi API.
// Include any external packages or services that will be required by these users.
type UsersApiService struct {
	creatingService creating.CreateUserService
	findAllService  finding.FindAllUsersUseCase
	findByIdService finding.FindUserByIdUseCase
}

// NewUsersApiService creates a default api users
func NewUsersApiService(creatingService creating.CreateUserService, findAllService finding.FindAllUsersUseCase, findByIdService finding.FindUserByIdUseCase) UsersApiServicer {
	return &UsersApiService{
		creatingService: creatingService,
		findAllService:  findAllService,
		findByIdService: findByIdService,
	}
}

// CreateUser - Save user into data storage
func (s *UsersApiService) CreateUser(ctx context.Context, dto UserDto) (ImplResponse, error) {
	err := s.creatingService.Create(ctx, dto.Name, dto.Firstname)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrInvalidUserID),
			errors.Is(err, users.ErrEmptyUserName),
			errors.Is(err, users.ErrEmptyFirstname):
			return Response(http.StatusNotFound, nil), err
		default:
			return Response(http.StatusInternalServerError, nil), err
		}
	}

	return Response(http.StatusCreated, nil), err
}

// GetAllUsers - Get all users
func (s *UsersApiService) GetAllUsers(ctx context.Context) (ImplResponse, error) {
	all, err := s.findAllService.FindAll(ctx)
	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return Response(http.StatusNotFound, nil), err
		default:
			return Response(http.StatusInternalServerError, nil), err
		}
	}

	return Response(http.StatusOK, UsersToUserDtos(all)), nil
}

// GetUserById - Get user by id
func (s *UsersApiService) GetUserById(ctx context.Context, userId string) (ImplResponse, error) {
	user, err := s.findByIdService.FindById(ctx, userId)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return Response(http.StatusNotFound, nil), err
		default:
			return Response(http.StatusInternalServerError, nil), err
		}
	}

	return Response(http.StatusOK, UserToUserDto(user)), nil
}
