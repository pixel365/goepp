package check

import (
	"encoding/xml"

	"github.com/pixel365/goepp/command"
)

func (c *Check) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	c.Domain = nil
	c.Contact = nil
	c.Host = nil

	var seen bool

	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		done, err := command.HandleToken(c, d, tok, &start, &seen, c.Name().String())
		if err != nil {
			return err
		}

		if done {
			return nil
		}
	}
}
