package email

import (
	"net/mail"
)

type adapter struct {
	address string
}

func New(address string) *adapter {
	if isValid := validateEmail(address); isValid {
		return &adapter{
			address,
		}
	}
	return &adapter{}
}

func validateEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func (e *adapter) Address() string {
	return e.address
}
