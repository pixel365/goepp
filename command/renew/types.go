package renew

import (
	"github.com/pixel365/goepp/command"
)

type Domain struct {
	Period *command.Period `xml:"period,omitempty"`
	command.DomainRef
	CurrentExpirationDate string `xml:"curExpDate"`
}
