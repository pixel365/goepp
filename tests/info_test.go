package tests

import (
	"testing"

	"github.com/pixel365/goepp"
)

func TestValidDomainInfo(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/info/domain/valid.xml")
}

func TestValidDomainInfoWithAuth(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/info/domain/valid_auth.xml")
}

func TestValidDomainInfoWithEmptyPassword(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/info/domain/invalid_auth.xml")
}

func TestValidContactInfo(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/info/contact/valid.xml")
}

func TestValidContactInfoWithAuth(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/info/contact/valid_auth.xml")
}

func TestInvalidContactInfoWithEmptyPassword(t *testing.T) {
	mustFailParse(t, &goepp.CmdParser{}, "testdata/info/contact/invalid_auth.xml")
}

func TestValidHostInfo(t *testing.T) {
	mustParseAndValidate(t, &goepp.CmdParser{}, "testdata/info/host/valid.xml")
}
