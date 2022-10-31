package server

import (
	users "api-template/internal"
	"reflect"
	"testing"
)

func Test_UserToUserDto(t *testing.T) {
	john, _ := users.NewUser("02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe")

	type args struct {
		user users.User
	}
	tests := []struct {
		name string
		args args
		want UserDto
	}{
		{
			name: "with invalid user",
			args: args{user: users.User{}},
			want: UserDto{},
		},
		{
			name: "with valid user",
			args: args{user: john},
			want: UserDto{
				Id:        "02b05d3e-43e7-4498-928f-e50a2eadde7b",
				Name:      "John",
				Firstname: "Doe",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserToUserDto(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserToUserDto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_UsersToUserDtos(t *testing.T) {
	john, _ := users.NewUser("02b05d3e-43e7-4498-928f-e50a2eadde7b", "John", "Doe")
	estelle, _ := users.NewUser("2c0c8411-bcf0-4d2d-8222-56e6a1cb6596", "Estelle", "Dawidowicz")

	type args struct {
		users []users.User
	}
	tests := []struct {
		name string
		args args
		want []UserDto
	}{
		{
			name: "with no users",
			args: args{users: nil},
			want: nil,
		},
		{
			name: "with one user",
			args: args{
				users: []users.User{john},
			},
			want: []UserDto{
				{
					Id:        "02b05d3e-43e7-4498-928f-e50a2eadde7b",
					Name:      "John",
					Firstname: "Doe",
				},
			},
		},
		{
			name: "with multiple users",
			args: args{
				users: []users.User{john, estelle},
			},
			want: []UserDto{
				{
					Id:        "02b05d3e-43e7-4498-928f-e50a2eadde7b",
					Name:      "John",
					Firstname: "Doe",
				},
				{
					Id:        "2c0c8411-bcf0-4d2d-8222-56e6a1cb6596",
					Name:      "Estelle",
					Firstname: "Dawidowicz",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UsersToUserDtos(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UsersToUserDtos() = %v, want %v", got, tt.want)
			}
		})
	}
}
