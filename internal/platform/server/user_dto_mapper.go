package server

import (
	users "github.com/rcebrian/users-service/internal"
)

// UserToUserDto maps a users.User to a server.UserDto
func UserToUserDto(user users.User) UserDto {
	var (
		id        = user.ID().String()
		name      = user.Name().String()
		firstname = user.Firstname().String()
	)

	return UserDto{
		Id:        &id,
		Name:      name,
		Firstname: firstname,
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
