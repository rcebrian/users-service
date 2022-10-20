package mysql

import (
	model "api-template/internal"
	"api-template/pkg/logger"
	"context"
	"database/sql"
)

// UserRepository is a MySQL mooc.UserRepository implementation.
type UserRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.UserRepository.
func NewCourseRepository(db *sql.DB) model.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Save(ctx context.Context, user model.User) error {
	panic("implement me")
}

func (u *UserRepository) FindById(ctx context.Context, id string) (model.User, error) {
	var user sqlUser

	err := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.id, &user.name, &user.firstname)
	if err != nil {
		return model.User{}, err
	}

	return model.NewUser(user.id, user.name, user.firstname), nil
}

func (u *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var user sqlUser
	var users []model.User

	defer rows.Close()
	for rows.Next() { //nolint:wsl
		if err = rows.Scan(&user.id, &user.name, &user.firstname); err != nil {
			logger.WithError(err).Error("error querying users")
			continue
		}

		users = append(users, model.NewUser(user.id, user.name, user.firstname))
	}

	if err != nil {
		logger.WithError(err).Error("error closing query")
	}

	return users, nil
}
