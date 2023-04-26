package server

import (
	"context"
	"errors"

	users "github.com/rcebrian/users-service/internal"
	"github.com/rcebrian/users-service/internal/users/creating"
	"github.com/rcebrian/users-service/internal/users/finding"
)

type UsersApiService struct {
	creatingService creating.CreateUserUseCase
	findAllService  finding.FindAllUsersUseCase
	findByIdService finding.FindUserByIdUseCase
}

var _ StrictServerInterface = (*UsersApiService)(nil)

// NewUsersApiServer creates a default api users
func NewUsersApiServer(creatingService creating.CreateUserUseCase, findAllService finding.FindAllUsersUseCase, findByIdService finding.FindUserByIdUseCase) StrictServerInterface {
	return &UsersApiService{
		creatingService: creatingService,
		findAllService:  findAllService,
		findByIdService: findByIdService,
	}
}

// GetAllUsers - Get all users
func (u UsersApiService) GetAllUsers(ctx context.Context, _ GetAllUsersRequestObject) (GetAllUsersResponseObject, error) {
	all, err := u.findAllService.FindAll(ctx)
	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return GetAllUsers404JSONResponse{}, err
		default:
			return GetAllUsers500JSONResponse{}, err
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
			return CreateUser400JSONResponse{UnsuccessfulResponseJSONResponse: UnsuccessfulResponseJSONResponse{
				Code:    nil,
				Message: &cause,
				Success: false,
			}}, err
		default:
			return CreateUser500JSONResponse{}, err
		}
	}

	return CreateUser201Response{}, nil
}

func (u UsersApiService) GetUserById(ctx context.Context, request GetUserByIdRequestObject) (GetUserByIdResponseObject, error) {
	user, err := u.findByIdService.FindById(ctx, request.UserId)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrNotFound):
			return GetUserById404JSONResponse{}, err
		default:
			return GetUserById500JSONResponse{}, err
		}
	}

	dto := UserToUserDto(user)

	return GetUserById200JSONResponse{Name: dto.Name, Firstname: dto.Firstname, Id: dto.Firstname}, err
}
