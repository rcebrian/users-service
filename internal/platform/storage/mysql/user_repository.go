package mysql

import (
	"context"
	"database/sql"

	users "github.com/rcebrian/users-service/internal"

	"github.com/sirupsen/logrus"

	"github.com/huandu/go-sqlbuilder"
)

// UserRepository is a MySQL users.UserRepository implementation.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a MySQL-based implementation of users.UserRepository.
func NewUserRepository(db *sql.DB) users.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save a users.User in persistence
func (r *UserRepository) Save(ctx context.Context, user users.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))

	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:        user.ID().String(),
		Name:      user.Name().String(),
		Firstname: user.Firstname().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// FindById search a users.User by unique id
func (r *UserRepository) FindById(ctx context.Context, id string) (user users.User, err error) {
	var dbUser sqlUser

	err = r.db.QueryRow("SELECT * FROM user WHERE id = ?  LIMIT 1", id).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Firstname)
	if err != nil {
		return users.User{}, err
	}

	user, err = users.NewUser(dbUser.ID, dbUser.Name, dbUser.Firstname)
	if err != nil {
		return users.User{}, nil
	}

	return user, nil
}

// FindAll get all users.User from persistence
func (r *UserRepository) FindAll(ctx context.Context) ([]users.User, error) {
	rows, err := r.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	var dbUser sqlUser
	var results []users.User

	defer rows.Close()
	for rows.Next() { //nolint:wsl
		if err = rows.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Firstname); err != nil {
			logrus.WithError(err).Error("error querying user")
			continue
		}

		user, newUserErr := users.NewUser(dbUser.ID, dbUser.Name, dbUser.Firstname)
		if err != nil {
			return nil, newUserErr
		}

		results = append(results, user)
	}

	if err != nil {
		logrus.WithError(err).Error("error closing query")
	}

	return results, nil
}
