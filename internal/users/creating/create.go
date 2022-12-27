package creating

import (
	"context"

	users "github.com/rcebrian/users-service/internal"

	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

//go:generate mockery --case=snake --outpkg=mocks --output=../mocks --name=CreateUserUseCase

type CreateUserUseCase interface {
	Create(ctx context.Context, name, firstname string) error
}

type createUserUseCase struct {
	repository users.UserRepository
}

func NewCreatingService(repository users.UserRepository) CreateUserUseCase {
	return createUserUseCase{repository: repository}
}

func (c createUserUseCase) Create(ctx context.Context, name, firstname string) error {
	id := uuid.New()

	user, err := users.NewUser(id.String(), name, firstname)
	if err != nil {
		logrus.WithError(err).Error("persisting user on database")
		return err
	}

	return c.repository.Save(ctx, user)
}
