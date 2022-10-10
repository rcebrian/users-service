package memory

import (
	model "api-template/internal"
	"context"
)

var users = []model.User{
	model.NewUser(1, "first", "test"),
	model.NewUser(2, "second", "test"),
	model.NewUser(3, "third", "test"),
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

func (u *UserRepository) FindById(_ context.Context, id int) (model.User, error) {
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
