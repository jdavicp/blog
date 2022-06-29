package author

import (
	ap "github.com/jdavicp/blog/internals/ports/left/author"
)

type User struct {
	name     string
	email    ap.EmailPort
	username ap.UsernamePort
	password string
}

func New(name string, email ap.EmailPort, username ap.UsernamePort, password string) *User {
	if email.Address() == "" || username.Username() == "" {
		return &User{}
	}
	return &User{
		name,
		email,
		username,
		password,
	}
}

func (a *User) Name() string {
	return a.name
}

func (a *User) Email() string {
	return a.email.Address()
}

func (a *User) Username() string {
	return a.username.Username()
}
