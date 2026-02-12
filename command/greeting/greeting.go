package greeting

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Greeting https://datatracker.ietf.org/doc/html/rfc5730#section-2.4
type Greeting struct {
	greeting GreetingInner
}

func (g Greeting) Marshal() ([]byte, error) {
	var b strings.Builder

	now := time.Now().UTC().Format(time.RFC3339)

	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	b.WriteString(`<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">`)
	b.WriteString(`<greeting>`)

	b.WriteString(`<svID>` + g.greeting.ServerID + `</svID>`)
	b.WriteString(`<svDate>` + now + `</svDate>`)

	b.WriteString(`<svcMenu>`)

	for _, v := range g.greeting.Versions {
		b.WriteString(`<version>` + v + `</version>`)
	}

	for _, lang := range g.greeting.Languages {
		b.WriteString(`<lang>` + lang.String() + `</lang>`)
	}

	for _, o := range g.greeting.ObjURI {
		b.WriteString(`<objURI>` + o.String() + `</objURI>`)
	}

	if len(g.greeting.Extensions) > 0 {
		b.WriteString(`<svcExtension>`)

		for _, ex := range g.greeting.Extensions {
			b.WriteString(`<extURI>` + ex.String() + `</extURI>`)
		}

		b.WriteString(`</svcExtension>`)
	}

	if g.greeting.Dcp != nil {
		g.greeting.Dcp.WriteXML(&b)
	}

	b.WriteString(`</svcMenu>`)

	b.WriteString(`</greeting>`)
	b.WriteString(`</epp>`)

	return []byte(b.String()), nil
}

func NewGreeting(greeting GreetingInner) Greeting {
	return Greeting{greeting}
}

type ObjURI string
type Extension string
type Language string

func (o ObjURI) String() string {
	return string(o)
}

func (o ObjURI) Validate() error {
	return nil
}

func (e Extension) String() string {
	return string(e)
}

func (e Extension) Validate() error {
	return nil
}

func (l *Language) String() string {
	return string(*l)
}

func (l *Language) Validate() error {
	if l == nil {
		return errors.New("language is nil")
	}

	lang := strings.ToLower(string(*l))
	if len(lang) != 2 {
		return errors.New("language must be 2 chars")
	}

	*l = Language(lang)

	return nil
}

// GreetingInner https://datatracker.ietf.org/doc/html/rfc5730#section-2.4
type GreetingInner struct {
	ServerID   string      `yaml:"sv_id"`
	Versions   []string    `yaml:"versions"`
	Extensions []Extension `yaml:"extensions"`
	Languages  []Language  `yaml:"languages"`
	Dcp        *Dcp        `yaml:"dcp,omitempty"`
	ObjURI     []ObjURI    `yaml:"objURI,omitempty"`
}

func (g GreetingInner) Validate() error {
	if g.ServerID == "" {
		return errors.New("server id is empty")
	}

	if len(g.Versions) == 0 {
		return errors.New("versions is empty")
	}

	for i, l := range g.Languages {
		if err := l.Validate(); err != nil {
			return fmt.Errorf("languages[%d] validation error: %w", i, err)
		}
	}

	if g.Dcp != nil {
		if err := g.Dcp.Validate(); err != nil {
			return err
		}
	}

	for i, uri := range g.ObjURI {
		if err := uri.Validate(); err != nil {
			return fmt.Errorf("objURI[%d] validation error: %w", i, err)
		}
	}

	return nil
}
