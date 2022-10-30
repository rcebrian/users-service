package creating

import (
	users "api-template/internal"
	"api-template/pkg/logger"
	"context"
)

type CreateUserService interface {
	Create(ctx context.Context, id, name, firstname string) error
}

type createUserUseCase struct {
	repository users.UserRepository
}

func NewCreatingService(repository users.UserRepository) CreateUserService {
	return createUserUseCase{repository: repository}
}

func (c createUserUseCase) Create(ctx context.Context, id, name, firstname string) error {
	user, err := users.NewUser(id, name, firstname)
	if err != nil {
		logger.WithError(err).Error("persisting user on database")
		return err
	}

	return c.repository.Save(ctx, user)
}
