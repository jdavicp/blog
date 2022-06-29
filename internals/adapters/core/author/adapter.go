package author

import (
	ap "github.com/jdavicp/blog/internals/ports/left/author"
)

type User struct {
	name     string
	email    ap.EmailPort
	username ap.UsernamePort
}

func New(name string, email ap.EmailPort, username ap.UsernamePort) *User {
	return &User{
		name,
		email,
		username,
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
