package update

import (
	"errors"

	"github.com/pixel365/goepp/command"
)

type Update struct {
	Domain  *DomainData
	Contact *ContactData
}

func (u *Update) Name() command.CommandName {
	return command.Update
}

func (u *Update) NeedAuth() bool {
	return true
}

func (u *Update) Validate() error {
	var notNil uint8

	if u.Domain != nil {
		notNil++
	}

	if u.Contact != nil {
		notNil++
	}

	if notNil != 1 {
		return errors.New("exactly one update command must be present")
	}

	return u.validate()
}

func (u *Update) validate() error {
	if u.Domain != nil {
		return u.Domain.Validate()
	}

	if u.Contact != nil {
		return u.Contact.Validate()
	}

	return nil
}
