package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidLogout(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/logout/valid.xml")
}

func TestInvalidLogout(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/logout/invalid.xml")
}
