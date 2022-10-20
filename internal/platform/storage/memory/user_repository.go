package memory

import (
	model "api-template/internal"
	"context"
)

var users = []model.User{
	model.NewUser("af03e53c-b847-4039-9691-6d7c8932e575", "first", "test"),
	model.NewUser("fb68413f-faa3-4e53-a04d-6e00407a313a", "second", "test"),
	model.NewUser("01c6a0a4-7f80-4111-a913-cda6d2aecc58", "third", "test"),
}

type UserRepository struct {
	users []model.User
}

func NewUserRepository() model.UserRepository {
	return &UserRepository{users: users}
}

func (u *UserRepository) Save(_ context.Context, user model.User) error {
	u.users = append(u.users, user)

	return nil
}

func (u *UserRepository) FindById(_ context.Context, id string) (model.User, error) {
	for i := range u.users {
		if u.users[i].ID() == id {
			return u.users[i], nil
		}
	}

	return model.User{}, nil
}

// FindAll implements the model.UserRepository interface.
func (u *UserRepository) FindAll(_ context.Context) ([]model.User, error) {
	return u.users, nil
}
