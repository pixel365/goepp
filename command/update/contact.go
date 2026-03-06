package update

import (
	"errors"
	"strings"

	"github.com/pixel365/goepp/command"
	"github.com/pixel365/goepp/command/create"
)

type ContactData struct {
	Add    *ContactAdd    `xml:"add,omitempty"`
	Remove *ContactRemove `xml:"rem,omitempty"`
	Change *ContactChange `xml:"chg,omitempty"`
	command.ContactRef
}

type ContactAdd struct {
	Statuses []ContactStatus `xml:"status,omitempty"`
}

type ContactRemove struct {
	Statuses []ContactStatus `xml:"status,omitempty"`
}

type ContactChange struct {
	Voice      *string             `xml:"voice,omitempty"`
	Fax        *string             `xml:"fax,omitempty"`
	Email      *string             `xml:"email,omitempty"`
	AuthInfo   *command.AuthInfo   `xml:"authInfo,omitempty"`
	Disclosure *Disclosure         `xml:"disclose,omitempty"`
	PostalInfo []create.PostalInfo `xml:"postalInfo,omitempty"`
}

type ContactStatus struct {
	Value string `xml:"s,attr"`
}

type Disclosure struct {
	Fields []DisclosureElement `xml:",any,omitempty"`
	Flag   bool                `xml:"flag,attr"`
}

type DisclosureElement struct {
	XMLName struct{} `xml:"-"`
	Type    string   `xml:"type,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// Validate https://datatracker.ietf.org/doc/html/rfc5733#section-3.2.5
func (c *ContactData) Validate() error {
	if c.ID == "" {
		return errors.New("contact:id is required")
	}

	if c.Add == nil && c.Remove == nil && c.Change == nil {
		return errors.New("contact:add, rem or chg is required")
	}

	if c.Add != nil {
		if err := c.Add.Validate(); err != nil {
			return err
		}
	}

	if c.Remove != nil {
		if err := c.Remove.Validate(); err != nil {
			return err
		}
	}

	if c.Change != nil {
		if err := c.Change.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (a *ContactAdd) Validate() error {
	if len(a.Statuses) == 0 {
		return errors.New("add:status is required")
	}

	for _, s := range a.Statuses {
		if err := s.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *ContactRemove) Validate() error {
	if len(r.Statuses) == 0 {
		return errors.New("rem:status is required")
	}

	for _, s := range r.Statuses {
		if err := s.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *ContactChange) Validate() error {
	if err := c.validateNotEmpty(); err != nil {
		return err
	}

	if err := c.validatePostalInfo(); err != nil {
		return err
	}

	if c.Voice != nil && strings.TrimSpace(*c.Voice) == "" {
		return errors.New("chg:voice is empty")
	}

	if c.Fax != nil && strings.TrimSpace(*c.Fax) == "" {
		return errors.New("chg:fax is empty")
	}

	if c.Email != nil && strings.TrimSpace(*c.Email) == "" {
		return errors.New("chg:email is empty")
	}

	if c.AuthInfo != nil {
		if err := c.AuthInfo.Validate(); err != nil {
			return err
		}
	}

	if c.Disclosure != nil {
		if err := c.Disclosure.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *ContactChange) validateNotEmpty() error {
	if len(c.PostalInfo) == 0 &&
		c.Voice == nil &&
		c.Fax == nil &&
		c.Email == nil &&
		c.AuthInfo == nil &&
		c.Disclosure == nil {
		return errors.New("chg:postalInfo, voice, fax, email, authInfo or disclose is required")
	}

	return nil
}

func (c *ContactChange) validatePostalInfo() error {
	for i := range c.PostalInfo {
		if err := c.PostalInfo[i].Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (s ContactStatus) Validate() error {
	if strings.TrimSpace(s.Value) == "" {
		return errors.New("status value is required")
	}

	return nil
}

func (d *Disclosure) Validate() error {
	if len(d.Fields) == 0 {
		return errors.New("disclose:name, org, addr, voice, fax or email is required")
	}

	return nil
}
