package goepp

import (
	"encoding/xml"
	"fmt"

	"github.com/pixel365/goepp/command"
)

type Parser interface {
	Parse([]byte) (command.Commander, error)
}

type CmdParser struct{}

func (c *CmdParser) Parse(payload []byte) (command.Commander, error) {
	var message EPP
	if err := xml.Unmarshal(payload, &message); err != nil {
		return nil, fmt.Errorf("unmarshal xml payload: %w", err)
	}

	if err := message.Validate(); err != nil {
		return nil, err
	}

	return message.Command()
}
