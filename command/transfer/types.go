package transfer

import (
	"encoding/xml"
	"errors"

	"github.com/pixel365/goepp/command"
)

type DomainTransfer struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	command.DomainRef
}

type ContactTransfer struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	command.ContactRef
}

type domainTransferXML struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	XMLName  xml.Name          `xml:"urn:ietf:params:xml:ns:domain-1.0 transfer"`
	Name     string            `xml:"name"`
}

type contactTransferXML struct {
	AuthInfo *command.AuthInfo `xml:"authInfo,omitempty"`
	XMLName  xml.Name          `xml:"urn:ietf:params:xml:ns:contact-1.0 transfer"`
	ID       string            `xml:"id"`
}

func (d *DomainTransfer) Validate(_ string) error {
	if d.AuthInfo != nil {
		if d.AuthInfo.Password == "" {
			return errors.New("domain:authInfo/domain:pw is required if authInfo is present")
		}
	}

	return d.DomainRef.Validate()
}

func (c *ContactTransfer) Validate(_ string) error {
	if c.AuthInfo != nil {
		if c.AuthInfo.Password == "" {
			return errors.New("contact:authInfo/contact:pw is required if authInfo is present")
		}
	}

	return c.ContactRef.Validate()
}
