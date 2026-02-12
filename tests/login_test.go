package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestLogin(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/login/valid.xml")
}
