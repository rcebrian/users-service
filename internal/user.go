package model

import "context"

// User is the data structure that represents a course.
type User struct {
	id        int
	name      string
	firstname string
}

// NewUser creates a new course.
func NewUser(id int, name string, firstname string) User {
	return User{
		id:        id,
		name:      name,
		firstname: firstname,
	}
}

// UserRepository defines the expected behaviour from a user storage.
type UserRepository interface {
	Save(ctx context.Context, user User) error
	FindById(ctx context.Context, id int) (User, error)
	FindAll(ctx context.Context) ([]User, error)
}

// ID returns the course unique identifier.
func (c User) ID() int {
	return c.id
}

// Name returns the course name.
func (c User) Name() string {
	return c.name
}

// Firstname returns the course firstname.
func (c User) Firstname() string {
	return c.firstname
}
