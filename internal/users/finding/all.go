package finding

import (
	users "api-template/internal"
	"context"
)

type FindAllUsersUseCase interface {
	FindAll(ctx context.Context) ([]users.User, error)
}

type useCase struct {
	repository users.UserRepository
}

func NewFindAllUsersUseCase(repository users.UserRepository) FindAllUsersUseCase {
	return useCase{repository: repository}
}

// FindAll get all users from database
func (s useCase) FindAll(ctx context.Context) ([]users.User, error) {
	return s.repository.FindAll(ctx)
}
