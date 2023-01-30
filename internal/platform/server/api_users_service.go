package server

import (
	"context"
	"errors"

	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/users/creating"
	"github.com/rcebrian/users-service/internal/users/finding"
)

var _ StrictServerInterface = (*UsersApiService)(nil)

type UsersApiService struct {
	creatingService creating.CreateUserUseCase
	findAllService  finding.FindAllUsersUseCase
	findByIdService finding.FindUserByIdUseCase
}

// GetAllUsers - Get all users
func (u UsersApiService) GetAllUsers(ctx context.Context, _ GetAllUsersRequestObject) (GetAllUsersResponseObject, error) {
	all, err := u.findAllService.FindAll(ctx)
	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return GetAllUsers404Response{}, err
		default:
			return GetAllUsers500Response{}, err
		}
	}

	allDto := UsersToUserDtos(all)

	return GetAllUsers200JSONResponse{Users: &allDto}, nil
}

func (u UsersApiService) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	err := u.creatingService.Create(ctx, *request.Body.Name, *request.Body.Firstname)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrInvalidUserID),
			errors.Is(err, users.ErrEmptyUserName),
			errors.Is(err, users.ErrEmptyFirstname):
			cause := err.Error()
			return CreateUser400JSONResponse{Errors: &cause}, err
		default:
			return CreateUser500Response{}, err
		}
	}

	return CreateUser201Response{}, nil
}

func (u UsersApiService) GetUserById(ctx context.Context, request GetUserByIdRequestObject) (GetUserByIdResponseObject, error) {
	user, err := u.findByIdService.FindById(ctx, request.UserId)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return GetUserById404Response{}, err
		default:
			return GetUserById500Response{}, err
		}
	}

	dto := UserToUserDto(user)

	return GetUserById200JSONResponse{Name: dto.Name, Firstname: dto.Firstname, Id: dto.Firstname}, err
}

// NewUsersApiServer creates a default api users
func NewUsersApiServer(creatingService creating.CreateUserUseCase, findAllService finding.FindAllUsersUseCase, findByIdService finding.FindUserByIdUseCase) StrictServerInterface {
	return &UsersApiService{
		creatingService: creatingService,
		findAllService:  findAllService,
		findByIdService: findByIdService,
	}
}