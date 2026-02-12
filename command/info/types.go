package info

import (
	"errors"

	"github.com/pixel365/goepp/command"
)

type DomainInfo struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	command.DomainRef
}

type ContactInfo struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	command.ContactRef
}

type HostInfo struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	command.HostRef
}

func (i *DomainInfo) Validate() error {
	if i.AuthInfo != nil {
		if i.AuthInfo.Password == "" {
			return errors.New("domain:authInfo/domain:pw is required if authInfo is present")
		}
	}

	return i.DomainRef.Validate()
}

func (i *ContactInfo) Validate() error {
	if i.AuthInfo != nil {
		if i.AuthInfo.Password == "" {
			return errors.New("contact:authInfo/contact:pw is required if authInfo is present")
		}
	}

	return i.ContactRef.Validate()
}

func (i *HostInfo) Validate() error {
	if i.AuthInfo != nil {
		if i.AuthInfo.Password == "" {
			return errors.New("host:authInfo/host:pw is required if authInfo is present")
		}
	}

	return i.HostRef.Validate()
}
