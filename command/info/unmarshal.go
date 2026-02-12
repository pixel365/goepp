package info

import (
	"encoding/xml"

	"github.com/pixel365/goepp/command"
)

func (i *Info) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	i.Domain = nil
	i.Contact = nil
	i.Host = nil

	var seen bool

	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		done, err := command.HandleToken(i, d, tok, &start, &seen, i.Name().String())
		if err != nil {
			return err
		}

		if done {
			return nil
		}
	}
}
