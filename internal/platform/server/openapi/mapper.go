package server

import (
	users "github.com/rcebrian/users-service/internal"
)

// UserToUserDto maps a users.User to a server.UserDto
func UserToUserDto(user users.User) UserDto {
	return UserDto{
		Id:        user.ID().String(),
		Name:      user.Name().String(),
		Firstname: user.Firstname().String(),
	}
}

// UsersToUserDtos maps an array of users.User to a server.UserDto array
func UsersToUserDtos(users []users.User) (usersDto []UserDto) {
	if users == nil {
		return nil
	}

	for i := range users {
		usersDto = append(usersDto, UserToUserDto(users[i]))
	}

	return usersDto
}
