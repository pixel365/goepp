package delete

import (
	"encoding/xml"

	"github.com/pixel365/goepp/command"
)

func (d *Delete) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	d.Domain = nil
	d.Contact = nil
	d.Host = nil

	var seen bool

	for {
		tok, err := dec.Token()
		if err != nil {
			return err
		}

		done, err := command.HandleToken(d, dec, tok, &start, &seen, d.Name().String())
		if err != nil {
			return err
		}

		if done {
			return nil
		}
	}
}
