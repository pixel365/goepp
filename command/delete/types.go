package delete

import (
	"github.com/pixel365/goepp/command"
)

type DomainDelete struct {
	command.DomainRef
}

type ContactDelete struct {
	command.ContactRef
}

type HostDelete struct {
	command.HostRef
}
