package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidHello(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/hello/valid.xml")
}

func TestInvalidHello(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/hello/invalid.xml")
}
