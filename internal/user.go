package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// UserID represents the course unique identifier.
type UserID struct {
	value string
}

// NewUserID instantiate the VO for UserID
func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

// String type converts the UserID into string.
func (id UserID) String() string {
	return id.value
}

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyUserName
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

// CourseFirstname represents the course duration.
type CourseFirstname struct {
	value string
}

func NewCourseFirstname(value string) (CourseFirstname, error) {
	if value == "" {
		return CourseFirstname{}, ErrEmptyFirstname
	}

	return CourseFirstname{
		value: value,
	}, nil
}

// String type converts the CourseFirstname into string.
func (duration CourseFirstname) String() string {
	return duration.value
}

// UserRepository defines the expected behaviour from a user storage.
type UserRepository interface {
	Save(ctx context.Context, user User) error
	FindById(ctx context.Context, id string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
}

// User is the data structure that represents a course.
type User struct {
	id        UserID
	name      CourseName
	firstname CourseFirstname
}

// NewUser creates a new course.
func NewUser(id, name, firstname string) (User, error) {
	idVO, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return User{}, err
	}

	firstnameVO, err := NewCourseFirstname(firstname)
	if err != nil {
		return User{}, err
	}

	return User{
		id:        idVO,
		name:      nameVO,
		firstname: firstnameVO,
	}, nil
}

// ID returns the course unique identifier.
func (c User) ID() UserID {
	return c.id
}

// Name returns the course name.
func (c User) Name() CourseName {
	return c.name
}

// Firstname returns the course duration.
func (c User) Firstname() CourseFirstname {
	return c.firstname
}
