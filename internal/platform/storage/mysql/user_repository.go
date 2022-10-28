package mysql

import (
	mooc "api-template/internal"
	"api-template/pkg/logger"
	"context"
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

// UserRepository is a MySQL mooc.UserRepository implementation.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a MySQL-based implementation of users.UserRepository.
func NewUserRepository(db *sql.DB) mooc.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user mooc.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))

	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:        user.ID().String(),
		Name:      user.Name().String(),
		Firstname: user.Firstname().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		logger.WithError(err).Error("error trying to persist user on database")
		return err
	}

	return nil
}

func (r *UserRepository) FindById(ctx context.Context, id string) (user mooc.User, err error) {
	var dbUser sqlUser

	err = r.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&dbUser.ID, &dbUser.Name, &dbUser.Firstname)
	if err != nil {
		return mooc.User{}, err
	}

	user, err = mooc.NewUser(dbUser.ID, dbUser.Name, dbUser.Firstname)
	if err != nil {
		return mooc.User{}, nil
	}

	return user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]mooc.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var dbUser sqlUser
	var users []mooc.User

	defer rows.Close()
	for rows.Next() { //nolint:wsl
		if err = rows.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Firstname); err != nil {
			logger.WithError(err).Error("error querying user")
			continue
		}

		user, newUserErr := mooc.NewUser(dbUser.ID, dbUser.Name, dbUser.Firstname)
		if err != nil {
			return nil, newUserErr
		}

		users = append(users, user)
	}

	if err != nil {
		logger.WithError(err).Error("error closing query")
	}

	return users, nil
}
