package finding

import (
	"context"
	"database/sql"
	"errors"

	users "github.com/rcebrian/users-service/internal"
)

//go:generate mockery --case=snake --outpkg=mocks --output=../mocks --name=FindUserByIdUseCase

type FindUserByIdUseCase interface {
	FindById(ctx context.Context, id string) (users.User, error)
}

type byIdUseCase struct {
	repository users.UserRepository
}

func NewFindUserByIdUseCase(repository users.UserRepository) FindUserByIdUseCase {
	return byIdUseCase{repository: repository}
}

// FindById get a user from database
func (s byIdUseCase) FindById(ctx context.Context, id string) (users.User, error) {
	user, err := s.repository.FindById(ctx, id)
	if err == nil {
		return user, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return users.User{}, users.ErrNotFound
	}

	return users.User{}, err
}
