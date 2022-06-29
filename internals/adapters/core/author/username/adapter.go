package username

import (
	"regexp"
)

type adapter struct {
	username string
}

func New(username string) *adapter {
	if isValid := validateUsername(username); isValid {
		return &adapter{
			username,
		}
	}
	return &adapter{}
}

func (a *adapter) Username() string {
	return a.username
}

func validateUsername(username string) bool {
	pat := "^[a-zA-Z0-9@#$]{5,8}$"
	reg, _ := regexp.Compile(pat)
	return reg.MatchString(username)
}
